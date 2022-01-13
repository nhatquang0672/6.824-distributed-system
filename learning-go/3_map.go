package main

import "strings"

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, v := range strings.Fields(s) {
		res[v] += 1
	}
	return res
}

func main() {
	// wc.Test(WordCount)
}
