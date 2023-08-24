package server

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func StartServer() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Server running")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("New Connection")
	var buffer = make([]byte, 1000)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buffer))
}
