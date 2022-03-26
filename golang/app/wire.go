//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
)

func CreateApp() *Application {
	wire.Build(
		AppProvider,
		CacheProvider,
		DatabaseProvider,
		ReportUseCaseProvider,
	)
	return new(Application)
}
