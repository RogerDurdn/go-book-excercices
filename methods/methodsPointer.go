package main

import "fmt"

// a method is a function that has a reciber of some type
// a type can be a struct or just a type
// a method can take a pointer or a value of a type

type Dog struct {
	name string
}

type Age int

func (dog *Dog) getName() string {
	return dog.name
}

//  if we use * this gonna take the pointer (reference to memory)
// instead if we dont use * they take the value (a copy of the memory - another)
func added(dog *Dog) {
	dog.name += " el perro"
}

func main() {
	p := Dog{"Juan"}
	fmt.Println(p.getName())
	added(&p)
	fmt.Println(p.getName())

}
