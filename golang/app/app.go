package app

import (
	"log"
	"net/http"

	"github.com/LuisDiazM/firestore-reports/domain"
	"github.com/LuisDiazM/firestore-reports/domain/usecases"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
)

type Application struct {
	Router          *mux.Router
	cacheGateway    domain.CacheGateway
	databaseGateway domain.DatabaseGateway
	reportUseCase   usecases.ReportsUsecase
}

func NewApplication(cacheGateway domain.CacheGateway, databaseGateway domain.DatabaseGateway, reportUseCase usecases.ReportsUsecase) *Application {
	return &Application{
		cacheGateway:    cacheGateway,
		databaseGateway: databaseGateway,
		reportUseCase:   reportUseCase,
	}
}

func (app Application) Start() {
	withGz := gziphandler.GzipHandler(app.Router)
	http.Handle("/", withGz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
