package main

import (
    "github.com/the-gigi/go-quote-service/cmd/service"
    "os"
)

func main() {
    port := 7777
    connectionString := os.Getenv("GO_QUOTE_SERVICE_CONNECTION_STRING")
    fmt.Println("")
    err := service.Run(port, connectionString)
    if err != nil {
        panic(err)
    }
}
