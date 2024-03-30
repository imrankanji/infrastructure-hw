package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://127.0.0.1:3000"
		// temp hack for testing allow true
		return true
	},
}

// Connect to Redis
var client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

func checkRedis() {
	// Ping Redis server to check the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)
}

func getBroadcasts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /broadcasts request\n")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	subs := client.Subscribe(ctx, "broadcasts")
	for {
		msg, err := subs.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}
		resp := []byte(msg.Payload)
		fmt.Println(msg.Payload)
		err = conn.WriteMessage(1, resp)
	}
}

func main() {

	checkRedis()

	http.HandleFunc("/broadcasts", getBroadcasts)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
