package main

import "fmt"

type Car struct {
	name  string
	speed int
}

func main() {
	m := make(map[string]int)
	m["hello"] = 10
	fmt.Println(m["hello"])
	fmt.Println(m)

	c := make(map[int]Car)
	c[1] = Car{name: "cheby"}
	c[2] = Car{name: "puk", speed: 100}

	fmt.Println(c[1])
	fmt.Println(c[2])
	fmt.Println(c)

	// in the literals the keys are required
	var ml = map[int]string{
		1: "pep",
	}
	fmt.Println(ml)

	elm, ok := ml[1]
	if elm == "" || !ok {
		println("this is bad thing")
	}

	// functions
	delete(ml, 1)

	elm, ok = ml[1]
	if !ok {
		fmt.Println("shuch a bad thing!")
	}

}
