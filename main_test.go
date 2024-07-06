package main

import (
    "bytes"
    "fmt"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    // Redirect standard output to buffer
    buf := new(bytes.Buffer)
    old := fmt.Swap(&buf)

    // Call the function that prints "Hello, World!"
    main()

    // Check the output
    expected := "Hello, World!\n"
    actual := buf.String()
    if actual != expected {
        t.Errorf("Expected: %s, got: %s", expected, actual)
    }

    // Restore original standard output
    fmt.Swap(old)
}

