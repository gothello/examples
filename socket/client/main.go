package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func ReadStdout(input chan string) {
	for {
		reader := bufio.NewReader(os.Stdout)
		d, err := reader.ReadString('\n')
		if err != nil {
			close(input)
			panic(err)
		}

		if d != "" {
			go WriteConn(input)
		}

	}
}

func WriteConn(out chan string) {

	var (
		conn net.Conn
		err  error
	)

	for {
		conn, err = net.Dial("tcp", ":8080")
		if err != nil {
			fmt.Println("reconnection")
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("status: connected")

		break
	}

	for {
		select {
		case b := <-out:
			_, err = conn.Write([]byte(b))

			if err != nil {
				log.Println(err)
			}
		default:
		}
	}

}

func main() {

	// conn, err := net.Dial("tcp", ":8080")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	input := make(chan string)

	go ReadStdout(input)

	<-make(chan bool)
}
