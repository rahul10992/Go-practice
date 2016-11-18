package main

import "golang.org/x/tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	var result = make(map[string]int)
	strArr := strings.Split(s, " ")
	for _, r := range strArr {
		c := string(r)
		result[c] = result[c] + 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
