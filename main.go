package main

import (
	"fmt"

	"github.com/shyam0507/go-message-bus/server"
)

func main() {
	fmt.Println("Starting the message bus")
	server.StartServer()
}
