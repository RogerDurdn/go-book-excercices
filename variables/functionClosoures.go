package main

import "fmt"

func main() {
	text, temp := apen(), apen()
	text("hello")
	fmt.Println(temp("who"))
	fmt.Println(text("how are"))
	fmt.Println(text("you"))
	fmt.Println(temp("you"))

}

func apen() func(string) string {
	wrd := ""
	return func(s string) string {
		wrd += " " + s
		return wrd
	}
}
