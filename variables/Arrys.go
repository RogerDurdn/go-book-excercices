package main

import "fmt"

func main() {
	var words [2]string
	words[0] = "uno"
	words[1] = "dos"
	fmt.Println(words[0], words[1])
	fmt.Println(words)

	numes := [2]int{1, 2}
	fmt.Println(numes[0])
	fmt.Println(numes)

	num := []int{1, 2, 3, 4, 5, 6}
	numba := num[1:4]

	fmt.Println(len(numes))
	fmt.Println(cap(numes))
	fmt.Println(cap(num))
	fmt.Println(cap(numba))

	var snil []int
	fmt.Println(len(snil), cap(snil), snil)

	if snil == nil {
		fmt.Println("this is nil")
	}

}
