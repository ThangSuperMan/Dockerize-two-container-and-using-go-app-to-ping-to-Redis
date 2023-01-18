package main

import (
	"PingRedis/folder"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func sum(a int, b int) int {
	if a > b {
		fmt.Println("a > b")
		return a
	} else {
		fmt.Println("a < b")
		return a
	}
}

func handleHomeRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleHomeRouter")
	fmt.Fprint(w, "home router")
}

func main() {
	fmt.Println("hello")
	client := redis.NewClient(&redis.Options{
		//container name:port
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	folder.CheckError(err)
	fmt.Println("Pong :>> ", pong)

	http.HandleFunc("/", handleHomeRouter)
	http.ListenAndServe(":3000", nil)
}
