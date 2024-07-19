package main

import "testing"

func TestSearch(t *testing.T) {
	directory := Directory{"test": "this is a simple test"}
	got := directory.Search("test")
	assertString(t, got, "this is a simple test")
}

func TestAdd(t *testing.T) {
	directory := Directory{}
	directory.Add("test", "this is a simple test")
	got := directory.Search("test")
	assertString(t, got, "this is a simple test")
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("get %v want %v", got, want)
	}
}
