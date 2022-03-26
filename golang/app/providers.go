package app

import (
	"github.com/LuisDiazM/firestore-reports/domain/usecases"
	"github.com/LuisDiazM/firestore-reports/infraestructure/cache"
	"github.com/LuisDiazM/firestore-reports/infraestructure/database"
	"github.com/google/wire"
)

var AppProvider = wire.NewSet(NewApplication)
var CacheProvider = wire.NewSet(cache.NewRedisCacheGatewayImpl)
var DatabaseProvider = wire.NewSet(database.NewDatabaseGatewayImp)
var ReportUseCaseProvider = wire.NewSet(usecases.NewReportUsecaseImp)
