// +build wireinject

package app

import (
	"github.com/google/wire"
)

func CreateApp() *Application {
	wire.Build(
		AppProvider,
		CacheProvider,
	)
	return new(Application)
}
