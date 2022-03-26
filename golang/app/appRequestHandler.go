package app

import (
	"net/http"

	"github.com/LuisDiazM/firestore-reports/domain/helpers"
)

const trxType = "useCases"

func (app Application) Index(response http.ResponseWriter, request *http.Request) {
	tx := helpers.Trace("index", trxType)
	app.reportUseCase.ExportDataUserResponses()
	defer tx.End()
	response.WriteHeader(http.StatusOK)
}
