package main

import "fmt"

func main() {
	primes := [3]int{1, 2, 3}
	var sl []int = primes[0:1]
	fmt.Println(sl)
	sl = append(sl, 3)
	fmt.Println(sl)
}
