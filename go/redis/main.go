package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	Connect()
}

func Connect() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "haigangliu",
		PoolSize: 1000,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

}