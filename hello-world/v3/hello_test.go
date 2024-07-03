package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("ning")
	want := "hello, ning"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
