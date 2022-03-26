package domain

import (
	"time"

	"github.com/LuisDiazM/firestore-reports/domain/models"
)

type DatabaseGateway interface {
	GetUsersResponses() (*[]models.UserResponsesDTO, error)
	Setup()
}

type CacheGateway interface {
	Get(prefix string, key string, data interface{}) (err error)
	Set(prefix string, key string, value interface{}, ttl time.Duration)
	Ping() (err error)
	Setup()
	Shutdown()
	Delete(prefix string, key string) (err error)
}
