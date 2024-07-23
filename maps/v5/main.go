package main

import (
	"errors"
	"fmt"
)

func main() {
	var ErrNotFound = errors.New("could not find the word you were looking for")
	fmt.Println(ErrNotFound == errors.New("could not find the word you were looking for")) // false

	fmt.Println(ErrNotFound == ErrNotFound) //true
}
