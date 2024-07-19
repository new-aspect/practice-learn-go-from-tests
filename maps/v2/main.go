package main

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

//type Dictionary struct {}

type Dictionary map[string]string

// Search 这里为什么不能使用 d *Dictionary, 这个时候如果改变d[str]的值，会不会在外部没有被改变，就像Dictionary的
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
