package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (app *application) InitRedis(name, port, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", name, port),
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}
