package redis

import (
	"encoding"
	"encoding/json"
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
	result, err := client.Set("as", "as1", 2*time.Second).Result()
	fmt.Println(result, err)
	result, err = client.Get("as").Result()
	fmt.Println(result, err)
	time.Sleep(2 * time.Second)
	result, err = client.Get("as").Result()
	fmt.Println(result, err)

	result, err = client.HMSet("a3", map[string]interface{}{"a1": "{\"name\":\"aaa\", \"age\":18}", "a2": "a22"}).Result()
	fmt.Println(result, err)
	var result1 = &User{}
	err = client.HGet("a3", "a1").Scan(result1)
	fmt.Println(result1, err)

}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var _ encoding.BinaryUnmarshaler = (*User)(nil)

func (user *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &user)
}
