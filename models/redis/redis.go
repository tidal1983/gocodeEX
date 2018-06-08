package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func OpenClient(addr, password string, DB int) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // "localhost:6379"
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}
