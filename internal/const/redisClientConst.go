package constants

const (
	Redis_client_file = `
package redis

import (
	"context"

	"github.com/go-redis/redis"

)
	
func NewClient(ctx context.Context) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "", // no password set
		DB:         0,  // use default DB)
		MaxRetries: 5,
	})
	return rdb
}
	
	`
)
