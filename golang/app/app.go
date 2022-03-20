package app

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"hibot.us/chatbots-report/domain"
)

type Application struct {
	Router       *mux.Router
	cacheGateway domain.CacheGateway
}

func NewApplication(cacheGateway domain.CacheGateway) *Application {
	return &Application{
		cacheGateway: cacheGateway,
	}
}

func (app Application) Start() {
	withGz := gziphandler.GzipHandler(app.Router)
	http.Handle("/", withGz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
