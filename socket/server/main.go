package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func reading(conn net.Conn) {
	for {

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {

			if err == io.EOF {
				log.Println("Connection closed.", conn.LocalAddr().String())
				return
			}

			log.Println(err)
			conn.Close()

			return
		}

		fmt.Print(conn.LocalAddr().String() + " received message:" + message)
		_, err = conn.Write([]byte("pong server" + message))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {

		conn, err := l.Accept()
		if err != nil {
			continue
		}

		go reading(conn)
	}

}
