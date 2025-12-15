package main

import "testing"

func TestGreeting(t *testing.T) {

	t.Run("saying hello to someone", func(t *testing.T) {
		got := greetingMessage("Anurag", "")
		want := "Hello Anurag"

		checkIfCorrect(got, want, t)
	})

	t.Run("saying hello to everyone", func(t *testing.T) {
		got := greetingMessage("", "")
		want := "Hello Guys"

		checkIfCorrect(got, want, t)
	})

	t.Run("In French", func(t *testing.T) {
		got := greetingMessage("Anurag", "French")
		want := "Bonjure Anurag"

		checkIfCorrect(got, want, t)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := greetingMessage("Anurag", "Spanish")
		want := "Hola Anurag"

		checkIfCorrect(got, want, t)
	})

	t.Run("In Hindi", func(t *testing.T) {
		got := greetingMessage("Anurag", "Hindi")
		want := "Namaste Anurag"

		checkIfCorrect(got, want, t)
	})
}

func checkIfCorrect(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
