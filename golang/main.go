package main

import (
	"github.com/LuisDiazM/firestore-reports/app"
	"github.com/gorilla/mux"
)

func main() {
	app := app.CreateApp()
	app.Router = mux.NewRouter()
	app.Setup()

	defer app.Shutdown()
	app.Start()
}
