package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(do("hello a you a mother"))
}

func do(s string) map[string]int {
	fiels := strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(fiels); i++ {
		val := 1
		if value, ok := m[fiels[i]]; ok {
			val = value + 1
		}
		m[fiels[i]] = val
	}
	return
}
