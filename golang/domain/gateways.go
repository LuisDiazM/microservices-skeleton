package domain

import "time"

type DatabaseGateway interface {
}

type CacheGateway interface {
	Get(prefix string, key string, data interface{}) (err error)
	Set(prefix string, key string, value interface{}, ttl time.Duration)
	Ping() (err error)
	Setup()
	Shutdown()
	Delete(prefix string, key string) (err error)
}
