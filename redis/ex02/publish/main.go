package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	ID       string `json:"id"`
	Username string `json:"user"`
	Pass     string `json:"pass"`
}

func (c *Client) Parser() ([]byte, error) {
	json, err := json.Marshal(c)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

type RedisPubSub struct {
	conn    *redis.Client
	channel string
}

func NewConnection(addr, pass string, db int, ch string) *RedisPubSub {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return &RedisPubSub{
		conn:    r,
		channel: ch,
	}
}

func (rds *RedisPubSub) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var client Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		fmt.Fprintf(w, http.StatusText(http.StatusFailedDependency))
		return
	}

	client.ID = uuid.NewString()

	body, err := client.Parser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = rds.conn.Publish(
		context.Background(),
		rds.channel,
		body,
	).Err()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "new client created id: %s", client.ID)

}

func main() {
	rb := NewConnection("localhost:6378", "admin", 0, "MY-CLIENT")
	if err := rb.conn.Ping(context.Background()).Err(); err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", rb.Handler)

	fmt.Println("Server on 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln(err)
	}
}
