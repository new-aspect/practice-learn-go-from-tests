package main

type Directory map[string]string

func (d Directory) Search(word string) string {
	return d[word]
}

func (d Directory) Add(work, content string) {
	d[work] = content
}
