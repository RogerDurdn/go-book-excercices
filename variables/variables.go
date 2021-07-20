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
	fmt.Printf(" %v  %T ", name, name)
	fmt.Println(sum(1, 2))
	prent("hello", sum(1, 2))
	fContinued("pep", sum(2, 2))
}

func sum(x, y int) int {
	return x + y
}

func prent(text string, num int) {
	for i := 0; i < num; i++ {
		fmt.Println(text)
	}
}

func fContinued(text string, num int) {
	ran := 3
	fmt.Println(ran)
	fmt.Println(num)
	for ran < num {
		fmt.Println(ran)
		fmt.Println(text)
		ran += ran
	}
}
