package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func haltOn(err error) {
	if err != nil {
		log.Fatal("Error here: ", err)
	}
}

func main() {
	fmt.Println("hi guys")
	client := redis.NewClient(&redis.Options{
		//container name:port
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	haltOn(err)
	fmt.Println("Pong :>> ", pong)
}
