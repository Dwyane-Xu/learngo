package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	var (
		lockKey    = "counter_lock"
		counterKey = "counter"
	)

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock restlt: ", lockSuccess)
		return
	}

	// counter++
	existResp := client.Exists(counterKey)
	existValue, err := existResp.Result()
	if err == nil && existValue == 0 {
		resp := client.Set(counterKey, 0, 0)
		_, err := resp.Result()
		if err != nil {
			println("init value error!")
		}
	}

	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	// unlock
	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
