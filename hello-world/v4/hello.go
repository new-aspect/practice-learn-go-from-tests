package main

import "fmt"

const englishHelloPrefix = "hello, "

func Hello(name string) string {
	return fmt.Sprintf("%s%s", englishHelloPrefix, name)
}
