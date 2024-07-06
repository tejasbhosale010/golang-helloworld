package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Hello, World!")
    
    // Keep the application running
    for {
        time.Sleep(time.Hour)
    }
}

