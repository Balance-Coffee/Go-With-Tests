package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say 'Hello, World' when an  empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello to person if a name is passed", func(t *testing.T) {
		got := Hello("Jake", "")
		want := "Hello, Jake"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jake", "Spanish")
		want := "Hola, Jake"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jake", "French")
		want := "Bonjour, Jake"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
