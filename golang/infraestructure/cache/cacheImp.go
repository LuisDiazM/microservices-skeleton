package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"hibot.us/chatbots-report/domain"
	"hibot.us/chatbots-report/domain/helpers"
)

type RedisCacheImpl struct {
	client       *redis.Ring
	cacheHandler *cache.Cache
	ctx          context.Context
}

func NewRedisCacheGatewayImpl() domain.CacheGateway {
	return &RedisCacheImpl{}
}

func (cacheMgt *RedisCacheImpl) Shutdown() {
	cacheMgt.client.Close()
}

func (cacheMgt *RedisCacheImpl) Setup() {
	cacheMgt.ctx = context.TODO()
	cacheMgt.client = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": os.Getenv("REDIS_HOST"),
		},
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	cacheMgt.cacheHandler = cache.New(&cache.Options{
		Redis:      cacheMgt.client,
		LocalCache: cache.NewTinyLFU(5, time.Minute),
	})

}

func (cacheMgt RedisCacheImpl) Set(prefix string, key string, value interface{}, ttl time.Duration) {
	tx := helpers.Trace("cache.set", "redis")
	defer tx.End()
	data, err := json.Marshal(value)
	if err == nil {
		_ = cacheMgt.cacheHandler.Set(&cache.Item{
			Ctx:   cacheMgt.ctx,
			Key:   prefix + ":" + key,
			Value: data,
			TTL:   ttl,
		})

	} else {
		helpers.TraceError(err, tx)
	}

}

func (cacheMgt RedisCacheImpl) Get(prefix string, key string, data interface{}) (err error) {
	tx := helpers.Trace("cache.get", "redis")
	defer tx.End()
	cachedHolder := make([]byte, 5)
	err = cacheMgt.cacheHandler.Get(cacheMgt.ctx, prefix+":"+key, &cachedHolder)
	if err == nil {
		_ = json.Unmarshal(cachedHolder, &data)
	} else {
		helpers.TraceError(err, tx)
	}
	return
}

func (cacheMgt RedisCacheImpl) Ping() (err error) {
	return cacheMgt.client.Ping(cacheMgt.ctx).Err()
}

func (cacheMgt RedisCacheImpl) Delete(prefix string, key string) (err error) {
	tx := helpers.Trace("cache.get", "redis")
	defer tx.End()
	err = cacheMgt.cacheHandler.Delete(cacheMgt.ctx, prefix+":"+key)
	return
}
