package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Holla, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Gaston", "French")
		want := "Bonjour, Gaston"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("João", "Portuguese")
		want := "Olá, João"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
