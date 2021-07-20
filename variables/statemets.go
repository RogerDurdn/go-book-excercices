package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrtt(x float64) string {
	if x < 0 {
		return "cat doit"
	}
	return "you cant"
}

func main() {
	fmt.Println(sqrtt(-2), sqrtt(2))
	fmt.Println(pow(2, 3, 10))
	fmt.Println(pow(3, 5, 10))
	osChoice()
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v)
	}
	return lim
}

func osChoice() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("is linux")
	case "windows":
		fmt.Println("is windows")
	default:
		fmt.Println("idknow")
	}
}
