package main

import (
	"fmt"
	"strings"
)

func main() {
	myslice := make([]int, 10)
	fmt.Println(len(myslice), cap(myslice))
	fmt.Println(myslice)
	myB := make([]string, 3)
	fmt.Println(myB)
	fmt.Println(len(myB), cap(myB))
	mA := myslice[:5]
	fmt.Println(mA)
	insideAnother()
}

func insideAnother() {
	table := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	table[0][0] = "O"
	table[0][1] = "X"
	table[0][2] = "X"
	table[0][0] = "O"
	table[2][0] = "X"

	fmt.Println(table)
	fmt.Println(" ")
	for i := 0; i < len(table); i++ {
		fmt.Printf("%s\n", strings.Join(table[i], " "))

	}
}
