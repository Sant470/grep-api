package v1

import (
	"log"
	"net/http"

	"github.com/sant470/grep-api/services"
	apptypes "github.com/sant470/grep-api/types"
	"github.com/sant470/grep-api/utils"
	"github.com/sant470/grep-api/utils/errors"
	"github.com/sant470/grep-api/utils/respond"
)

type SearchHandler struct {
	lgr *log.Logger
	svc *services.SearchService
}

func NewSearchHandler(lgr *log.Logger, svc *services.SearchService) *SearchHandler {
	return &SearchHandler{lgr, svc}
}

func (sh *SearchHandler) Search(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	var req apptypes.SearchReq
	if err := utils.Decode(r, &req); err != nil {
		return err
	}
	if req.SearchKeyword == "" || req.From > req.To {
		return errors.BadRequest("invalid params")
	}
	result, err := sh.svc.Search(&req)
	if err != nil {
		return err
	}
	return respond.OK(rw, result)
}
