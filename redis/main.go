package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.115.203.249:6379",
		Password: "",
		DB:       0,
	})
}
func main() {
	//ctx := context.Background()

	err := rdb.Set("say", "hello", 0).Err()
	if err != nil {
		panic(err)
	}
	value, err := rdb.Get("say").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("say :", value)
}
