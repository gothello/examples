package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func StartConnection() net.Conn {

	var (
		conn net.Conn
		err  error
	)
	for {
		conn, err = net.Dial("tcp", ":8080")
		if err == nil {
			fmt.Println("connected")
			break
		}

		fmt.Println("connecting.................")
		time.Sleep(time.Second)
	}

	return conn
}

func ReadStdout(input chan string) {
	for {
		reader := bufio.NewReader(os.Stdout)
		d, err := reader.ReadString('\n')
		if err != nil {
			close(input)
			panic(err)
		}

		if d != "" {
			input <- d
		}

	}
}

func ReadInputServer(conn net.Conn, out chan string) {

	body := make([]byte, 1024)
	for {
		b, err := conn.Read(body)
		if err != nil {
			log.Println(err)
			return
		}

		out <- string(body[:b])
	}
}

func main() {

	input := make(chan string)
	output := make(chan string)

START:
	for {
		conn := StartConnection()

		go ReadStdout(input)
		go ReadInputServer(conn, output)

		for {
			select {
			case i := <-input:
				_, err := conn.Write([]byte(i))
				if err != nil {
					log.Println(err)
					conn.Close()
					continue START
				}
			case o := <-output:
				fmt.Printf("Received message: %s", o)
			}
		}
	}
}
