package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello, World!")
    // Keep the container running indefinitely
    for {
        time.Sleep(time.Second)
    }
}

