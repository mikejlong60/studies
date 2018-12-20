package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)

	m := make(map[string]int)

	for i := range fields {
		word := m[fields[i]] + 1
		m[fields[i]] = word
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
