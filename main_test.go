package main

import (
    "bytes"
    "fmt"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    var buf bytes.Buffer
    fmt.Fprint(&buf, "Hello, World!")

    expected := "Hello, World!"
    if buf.String() != expected {
        t.Errorf("expected %s but got %s", expected, buf.String())
    }
}

