package app

import (
	"github.com/google/wire"
	"hibot.us/chatbots-report/infraestructure/cache"
)

var AppProvider = wire.NewSet(NewApplication)
var CacheProvider = wire.NewSet(cache.NewRedisCacheGatewayImpl)
