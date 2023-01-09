// https://go.dev/tour/moretypes/23
package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	substrings := strings.Fields(s)
	m := map[string]int{}
	for _, str := range substrings {
		m[str] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
