package main

import "fmt"

const (
	helloPrefix = "Hell, "
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s%s", helloPrefix, name)
}
