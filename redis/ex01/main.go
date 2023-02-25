package main

import (
	"context"
	"fmt"
	"log"
	"time"

	redis "github.com/redis/go-redis/v9"
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

	rd := Conn("localhost:6378", "admin", 0)

	timeout := time.Duration(100) * time.Millisecond

	err := rd.Set(ctx, "example-redis", "Hello Redis", timeout).Err()
	if err != nil {
		log.Fatalln(err)
	}

	// err = rd.Append(ctx, "example-redis-01", "redis is beautiful").Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	value, err := rd.Get(ctx, "example-redis").Result()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("example-redis:%s\n", value)

}
