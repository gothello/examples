package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type InputRedis struct {
	ID       string `json:"id"`
	Username string `json:"user"`
	Pass     string `json:"pass"`
}

func NewConn(addr, pass string, db int) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	return r
}

func main() {
	r := NewConn("localhost:6378", "admin", 0)

	sub := r.Subscribe(context.Background(), "MY-CLIENT")

	fmt.Println("Start rabbitmq reading")

	for {
		m, err := sub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		var data InputRedis

		if err := json.Unmarshal([]byte(m.Payload), &data); err != nil {
			log.Println(err)
		}

		fmt.Println(data)

	}
}
