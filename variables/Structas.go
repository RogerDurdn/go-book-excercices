package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	yo := Person{}
	yo.name = "juan"
	fmt.Println(yo.name)
	yo.name = "pepe"
	fmt.Println(yo.name)
	fmt.Println(yo)

	yoA := Person{name: "Pale"}
	fmt.Println(yoA)
	yoB := &Person{name: "joselo"}
	println(yoB.name)
	yoB.age = 20
	println(yoB.age)

}
