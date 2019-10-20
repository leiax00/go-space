package redis

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/go-redis/redis/v7"

func TestSetFunc(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root", // no password set
		DB:       0,      // use default DB
	})
	result, err := client.Set("as", "as1", 10*time.Second).Result()
	fmt.Println(result, err)
	result, err = client.Get("as").Result()
	fmt.Println(result, err)
	time.Sleep(10 * time.Second)
	result, err = client.Get("as").Result()
	fmt.Println(result, err)
}
