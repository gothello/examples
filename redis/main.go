package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func Conn(addr, pass string, db int) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return r
}

func main() {
	ctx := context.Background()

	rd := Conn("localhost:6379", "", 0)

	err := rd.Set(ctx, "example-redis", "Hello Redis", 0).Err()
	if err != nil {
		log.Fatalln(err)
	}

	value, err := rd.Get(ctx, "example-redis").Result()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("example-redis:%s\n", value)

}
