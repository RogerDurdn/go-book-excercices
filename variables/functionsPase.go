package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(get("wahts app mothe", pas))
}

func get(text string, fnp func(string) []string) []string {
	wrds := fnp(text)
	for i := 0; i < len(wrds); i++ {
		wrds[i] = strings.ToUpper(wrds[i])
	}
	return wrds
}

func pas(s string) []string {
	s += " hello there is me"
	return strings.Fields(s)
}
