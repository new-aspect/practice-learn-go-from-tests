package main

type Dictionary map[string]string

func (d Dictionary) Add(key, content string) {
	d[key] = content
}

func (d Dictionary) Search(key string) (string, error) {
	return d[key], nil
}
