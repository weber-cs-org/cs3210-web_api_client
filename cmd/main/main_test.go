package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func captureLogOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	f()

	return buf.String()
}

func TestCanary(t *testing.T) {

}

func TestOutput(t *testing.T) {
	// Arrange
	expected := ""
	// Act
	actual := captureLogOutput(main)
	// Assert
	if !strings.Contains(actual, expected) {
		t.Errorf("Got '%s' but does not contain expected '%s'",
			actual,
			expected)
	}
}
