package main

import (
    "fmt"
    "log"
    "net/http"
)

 func main() {
    fmt.Printf("Starting web app at ioppppp 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}