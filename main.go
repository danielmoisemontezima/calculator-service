package main

import (
    "log"
    "net/http"

    "github.com/danielmoisemontezima/calculator-service/handlers"
)

func main() {
    http.HandleFunc("/calculate", handlers.Calculate)
    log.Println("Calculator service running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
