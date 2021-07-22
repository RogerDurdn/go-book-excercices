package main

import "fmt"

func main() {
	fib := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func fibo() func() int {
	m := make(map[int]int)
	inx := 0
	return func() int {
		var curr int
		if len(m) < 2 {
			curr = inx
		} else {
			curr = m[inx-1] + m[inx-2]
		}
		m[inx] = curr
		inx++
		return curr
	}
}
