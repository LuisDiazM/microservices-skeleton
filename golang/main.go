package main

import (
	"github.com/gorilla/mux"
	"hibot.us/chatbots-report/app"
)

func main() {
	app := app.CreateApp()
	app.Router = mux.NewRouter()
	app.Setup()

	defer app.Shutdown()
	app.Start()
}
