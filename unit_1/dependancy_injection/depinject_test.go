package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("printing works correctly", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(&buffer, "Anurag")

		got := buffer.String()
		want := "Hello, Anurag"

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})
}
