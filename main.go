package main

import (
	"log"
	"net/http"

	"github.com/sant470/grep-api/apis"
	handlers "github.com/sant470/grep-api/apis/v1"
	"github.com/sant470/grep-api/config"
	"github.com/sant470/grep-api/services"
)

func main() {
	lgr := log.Default()
	lgr.Println("info: starting the server")
	router := config.InitRouters()
	searchSvc := services.NewSearchService(lgr)
	searchHlr := handlers.NewSearchHandler(lgr, searchSvc)
	apis.InitSerachRoutes(router, searchHlr)
	http.ListenAndServe("localhost:8000", router)
}
