package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("用中文", func(t *testing.T) {
		got := Hello("ning", "Chinese")
		want := "你好, ning"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
