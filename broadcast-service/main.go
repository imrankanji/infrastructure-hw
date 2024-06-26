package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	for {
		err := client.Publish(ctx, "broadcasts", "Hello world").Err()
		if err != nil {
			log.Fatal(err)
		}
		r := rand.Intn(10)
		fmt.Println(r)
		time.Sleep(time.Duration(r) * time.Second)

	}
}
