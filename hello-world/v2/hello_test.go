package main

import "testing"

func TestSayHello(t *testing.T) {
	got := SayHello()
	want := "hello, world"

	if got != want {
		t.Errorf("want get %s, but get %s", got, want)
	}
}
