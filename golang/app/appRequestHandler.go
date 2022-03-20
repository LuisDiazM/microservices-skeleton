package app

import (
	"net/http"

	"hibot.us/chatbots-report/domain/helpers"
)

const trxType = "useCases"

func (app Application) Index(response http.ResponseWriter, request *http.Request) {
	tx := helpers.Trace("index", trxType)
	defer tx.End()
	response.WriteHeader(http.StatusOK)
}
