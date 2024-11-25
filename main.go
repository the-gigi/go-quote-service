package main

import (
	"fmt"
	"github.com/the-gigi/go-quote-service/cmd/service"
	"os"
	"strconv"
	"strings"
)

func main() {
	var port = 7777
	servicePort := os.Getenv("GO_QUOTE_SERVICE_PORT")
	var err error
	if servicePort != "" {
		parts := strings.Split(servicePort, ":")
		if len(parts) > 2 {
			fmt.Println("Invalid port string")
			os.Exit(1)
		}

		port, err = strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	connectionString := os.Getenv("GO_QUOTE_SERVICE_CONNECTION_STRING")
	fmt.Println("Listening on port ", port)
	err = service.Run(port, connectionString)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
