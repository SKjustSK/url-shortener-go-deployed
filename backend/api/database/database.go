package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// global context used for all Redis operations
var Ctx = context.Background()

// initializes a connection to a specific Redis database (0-15)
func CreateClient(dbNo int) *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	// Upstash only allows db 0 to be used, therefore opting db 0 always
	opt.DB = 0

	rdb := redis.NewClient(opt)

	// ADD THIS: Test the connection immediately
	if err := rdb.Ping(Ctx).Err(); err != nil {
		// This will print the EXACT error to your docker logs
		panic("Redis connection failed: " + err.Error())
	}

	return rdb
}
