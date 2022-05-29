package main

import (
	"fmt"
	"github.com/the-gigi/go-quote-service/cmd/service"
	"os"
)

const (
	port int = 7777
)

func main() {
	connectionString := os.Getenv("GO_QUOTE_SERVICE_CONNECTION_STRING")
	fmt.Println("Listening on port ", port)
	err := service.Run(port, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
