package main

import (
	"fmt"
	"github.com/the-gigi/go-quote-service/cmd/service"
	"os"
	"strconv"
)

func main() {
	servicePort := os.Getenv("GO_QUOTE_SERVICE_PORT")
	if servicePort == "" {
		servicePort = "80"
	}

	port, err := strconv.Atoi(servicePort)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	connectionString := os.Getenv("GO_QUOTE_SERVICE_CONNECTION_STRING")
	fmt.Println("Listening on port ", port)
	err = service.Run(port, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
