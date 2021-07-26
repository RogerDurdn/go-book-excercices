package main

import "fmt"

func suma(numbs []int, c chan int) {
	sum := 0
	for _, v := range numbs {
		sum += v
	}
	c <- sum
}

func main() {
	numbs := []int{1, 2, 3, 10, 10, 10}
	c := make(chan int)
	go suma(numbs[:len(numbs)/2], c)
	go suma(numbs[len(numbs)/2:], c)

	r1, r2 := <-c, <-c

	fmt.Printf("%v , %v, result:%v", r1, r2, r1+r2)
}
