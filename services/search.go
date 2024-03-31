package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sant470/grep-api/lib/s3"
	apptypes "github.com/sant470/grep-api/types"
	"github.com/sant470/grep-api/utils/errors"
)

type SearchService struct {
	lgr *log.Logger
}

func NewSearchService(lgr *log.Logger) *SearchService {
	return &SearchService{lgr}
}

func (searchSvc *SearchService) Search(searchReq *apptypes.SearchReq) (*apptypes.SearchResult, *errors.AppError) {
	var searchResult apptypes.SearchResult
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	/*
		Limiting network connections at a time, should be configurable
		Ideally, semaphore should be on combination of network connections and memory in use
	*/
	sem := make(chan struct{}, 10)

	// quit will force the goroutines to stop what they are doing and return
	quit := make(chan struct{})

	// read the matches from the matchChannel
	matchChannel := make(chan apptypes.Match, 100)

	// this is to ensure that request returns within given time frame
	go func() {
		<-ctx.Done()
		close(quit)
		close(matchChannel)
	}()

	cli, err := s3.NewClient(searchSvc.lgr)
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}
	for start := searchReq.From; start < searchReq.To; {
		hour, date := dateTimeFormat(start)
		pathSuffix := fmt.Sprintf("%s%s%s%s%s", date, "/", hour, ".", "txt")
		go func(suffix string) {
			sem <- struct{}{}
			searchSvc.searchRemoteFile(cli, suffix, searchReq.SearchKeyword, matchChannel, quit)
			<-sem
		}(pathSuffix)
		start += int64(3600)
	}
	for match := range matchChannel {
		searchResult.Matches = append(searchResult.Matches, match)
	}
	return &searchResult, nil
}
