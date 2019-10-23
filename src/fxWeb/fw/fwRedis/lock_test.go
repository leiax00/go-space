package fwRedis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

var lockName = "lock:fwRedis"
var acquireTimeOut = 9 * time.Minute
var lockTimeOut = 4 * time.Second

func TestLock(t *testing.T) {
	count := 100
	var wg sync.WaitGroup
	fwRedisClient = getClient()
	for count > 0 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			s, e := GetLock(lockName, acquireTimeOut, lockTimeOut)
			fmt.Println(time.Now(), id, s, e)
			duration, _ := time.ParseDuration(strconv.Itoa(rand.Intn(4)) + "s")
			fmt.Println(time.Now(), id, "sleep:", duration.String())
			time.Sleep(duration)
			r := ReleaseLock(lockName, s)
			fmt.Println(time.Now(), id, s, e, "sleep:", duration.String(), r)
		}(count)
		count--
	}
	wg.Wait()
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "root",
		DialTimeout:  time.Minute,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	})
}
