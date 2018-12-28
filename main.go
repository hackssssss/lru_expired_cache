package main

import (
	"fmt"
	"lru_expiredmap/lru_expiredcache"
	"strconv"
	"time"
)

func main() {
	cache := lru_expiredcache.NewCache(3)

	cache.SetWithExpiredTime(2, []byte("2"), 2)
	time.Sleep(time.Second * 1)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Exist(2))
	time.Sleep(time.Second)
	fmt.Println(cache.Exist(2))

	for i := 0; i < 1000; i++ {
		cache.SetWithExpiredTime(i, []byte(strconv.Itoa(i)), 1)
	}
	cache.DebugShowLruList()
	cache.DebugShowMapData()
	fmt.Print(cache.Get(997))
	cache.DebugShowLruList()
	time.Sleep(time.Second)
	fmt.Println(cache.Get(998))
	cache.DebugShowLruList()
}
