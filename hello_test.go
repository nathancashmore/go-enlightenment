package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in English if language unknown", func(t *testing.T) {
		got := Hello("Chris", "Unknown")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
}
