package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	b := make([]byte, 1024)

	for {

		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		_, err = conn.Read(b)
		if err != nil {
			log.Println(err)
		}

		if string(b) != "" {
			fmt.Print(string(b))
		}
	}

}
