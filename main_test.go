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
	old := os.Stdout // keep backup of the real stdout
	os.Stdout = &buf

	// Restore the original stdout once we are done
	defer func() {
		os.Stdout = old
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

