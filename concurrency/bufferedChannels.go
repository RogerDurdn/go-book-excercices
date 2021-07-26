package main

import "fmt"

func main() {
	chBuffer := make(chan int, 2)
	chBuffer <- 22
	chBuffer <- 2
	chBuffer <- 234
	fmt.Println(<-chBuffer)
	fmt.Println(<-chBuffer)

}
