package main

import "fmt"

// this is a multiple declaration
var (
	name string  = "roger"
	age  int     = 28
	size float32 = 10.9
)

func main() {
	cthing := "Hello"
	fmt.Println(cthing)
	fmt.Printf(" %v , %t", name, name)
}
