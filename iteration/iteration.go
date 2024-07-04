package main

import "strings"

func Repeat(a string, iterate int) (repeat string) {
	for i := 0; i < iterate; i++ {
		repeat += a
	}
	return
}

func RepeatFaster(s string, iterate int) (repeat string) {
	var builder strings.Builder
	builder.Grow(len(s) * iterate)
	for i := 0; i < iterate; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}
