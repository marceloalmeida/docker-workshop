package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	for {
		ping(client)
		setGet(client)
		incr(client)
		time.Sleep(1000 * time.Millisecond)
	}
}

func incr(client *redis.Client) {
	result, err := client.Incr("counter").Result()
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}

	fmt.Println(result)
}

func ping(client *redis.Client) {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func setGet(client *redis.Client) {
	err := client.Set("key", "docker-workshop", 0).Err()
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		//panic(err)
		fmt.Println(err)
	} else {
		fmt.Println("key2", val2)
	}
}
