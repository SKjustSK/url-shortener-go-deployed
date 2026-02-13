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
	// Parse the connection string from the environment variable
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		// If the URL is invalid, we can't proceed.
		// Since the function signature only returns *redis.Client, we panic here.
		panic(err)
	}

	// Override the DB number with the argument passed to the function
	opt.DB = dbNo

	// Create the client with the parsed options
	rdb := redis.NewClient(opt)

	return rdb
}
