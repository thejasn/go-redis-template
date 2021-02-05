package main

import (
	"fmt"

	"github.com/thejasn/go-redis-template/pkg/redis"
)

func main() {
	fmt.Println("In main")
	client := redis.NewClient().Cluster
	err := client.Set("keyaa", "thejas", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("keyaa").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("keyaa", val)
}
