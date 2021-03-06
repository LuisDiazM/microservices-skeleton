package app

import (
	"net/http"
)

func enableCors(responseWriter *http.ResponseWriter) {
	//Enable CORS
	(*responseWriter).Header().Set("Access-Control-Allow-Credentials", "true")
	(*responseWriter).Header().Set("Access-Control-Allow-Methods", "*")
	(*responseWriter).Header().Set("Access-Control-Allow-Origin", "*")
	(*responseWriter).Header().Set("Access-Control-Allow-Headers", "*")
	(*responseWriter).Header().Set("Access-Control-Max-Age", "3600")
}

func (app Application) Setup() {
	app.cacheGateway.Setup()
	app.databaseGateway.Setup()
	app.Router.Use(app.route)

	reportsV1 := app.Router.PathPrefix("/api/v1").Subrouter()
	reportsV1.Path("/").HandlerFunc(app.Index).Methods(http.MethodGet, http.MethodOptions)

}

func (app Application) route(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		enableCors(&responseWriter)
		if (*request).Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(responseWriter, request)
	})
}

func (app Application) Shutdown() {
	app.cacheGateway.Shutdown()
}
