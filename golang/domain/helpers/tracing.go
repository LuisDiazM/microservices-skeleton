package helpers

import (
	"fmt"
	"log"
	"time"

	"go.elastic.co/apm"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func Trace(trx string, trxType string) *apm.Transaction {
	return apm.DefaultTracer.StartTransaction(trx, trxType)
}

func TraceError(err error, trx *apm.Transaction) {
	e := apm.DefaultTracer.NewError(err)
	e.SetTransaction(trx)
	trx.Result = fmt.Sprintf("Error: %s", trx.Name)
}

func Span(trx *apm.Transaction, spanName string, spanType string) *apm.Span {
	return trx.StartSpan(spanName, spanType, nil)
}
