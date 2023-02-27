package cache

import (
	"encoding/json"
	"fmt"
	"karthikeyan/books/model"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) BookCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key string, value string) {
	client := cache.getClient()

	// json, err := json.Marshal(value)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(value)

	status := client.Set(key, value, cache.expires*time.Hour)
	fmt.Println()
	fmt.Println(status)

}

func (cache *RedisCache) Get(key string) string {

	client := cache.getClient()

	val, err := client.Get(key).Result()

	fmt.Println(val)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var book model.Book

	err = json.Unmarshal([]byte(val), &book)
	if err != nil {
		panic(err)
	}

	return book.Title

}
