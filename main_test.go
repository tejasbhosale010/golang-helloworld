package main

import (
    "bytes"
    "fmt"
    "os"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    // Redirect stdout to a buffer
    var buf bytes.Buffer
    old := os.Stdout
    os.Stdout = &buf
    defer func() {
        os.Stdout = old // Restore original stdout
    }()

    // Call the function that prints "Hello, World!"
    main()

    // Check the output
    expected := "Hello, World!\n"
    actual := buf.String()
    if actual != expected {
        t.Errorf("Expected: %s, got: %s", expected, actual)
    }
}

