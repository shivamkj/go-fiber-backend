package helper

import (
	"context"

	"github.com/qnify/api-server/utils/errors"
	"github.com/redis/go-redis/v9"
)

var redisInit = false

func GetRedis(redisConnURL string) *redis.Client {
	if redisInit {
		panic(errors.New("redis is already initialised"))
	}

	opt, err := redis.ParseURL(redisConnURL)
	if err != nil {
		panic(errors.Wrap("invalid redis connection url", err))
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.TODO()).Result()
	if err != nil {
		panic(errors.Wrap("couldn't connect to redis", err))
	}

	redisInit = true
	return client
}
