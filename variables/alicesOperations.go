package main

import "fmt"

func main() {
	var maslice []int
	fmt.Println(maslice)

	maslice = append(maslice, 1, 2, 3, 4)
	fmt.Println(maslice)

	maslice = append(maslice, 23, 1, 2, 3)
	fmt.Println(maslice)
	fmt.Println(len(maslice), cap(maslice), maslice)
	fmt.Println("the for range")
	for i, v := range maslice {
		fmt.Println(i, v)
	}
	fmt.Println("continue to for range")
	for _, v := range maslice {
		fmt.Println(v)
	}
}
