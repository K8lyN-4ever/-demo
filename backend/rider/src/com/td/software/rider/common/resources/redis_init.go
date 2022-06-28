package resources

import (
	"github.com/go-redis/redis"
)

var RedisDb *redis.Client

func InitClient() (err error) {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	_, err = RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
