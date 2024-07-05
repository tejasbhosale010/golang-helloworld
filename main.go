package main

import (
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Setup HTTP handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // Start HTTP server in a goroutine
    go func() {
        fmt.Println("Server listening on port 8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Printf("Failed to start server: %v\n", err)
            os.Exit(1) // Exit if HTTP server fails to start
        }
    }()

    // Setup graceful shutdown
    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
    <-shutdown

    fmt.Println("Shutting down server...")

    // Perform cleanup tasks if necessary

    fmt.Println("Server stopped gracefully")
}

