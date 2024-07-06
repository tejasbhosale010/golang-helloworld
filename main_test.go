package main

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
)

func TestHelloWorldProgram(t *testing.T) {
    // Run the Docker container
    cmd := exec.Command("docker", "run", "--rm", "go-hello-world")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        t.Fatalf("Failed to run Docker container: %v", err)
    }

    // Check if "Hello, World!" is present in the output
    output := out.String()
    if !strings.Contains(output, "Hello, World!") {
        t.Errorf("Expected output 'Hello, World!', got: %s", output)
    }

    fmt.Printf("Output: %s", output)
}

