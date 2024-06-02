package main

import "testing"

func TestHello(t *testing.T) {
	// got := Hello("Chris")
	// want := "Hello, Chris"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Mandy", "Spanish")
		want := "Hola, Mandy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mandy", "French")
		want := "Bonjour, Mandy"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
