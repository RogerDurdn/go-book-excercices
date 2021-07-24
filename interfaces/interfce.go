package main

import "fmt"

/*
The interface is used explicitly just by implements the
declared methods on there, this is same as methods by the
declard reciber (value/pointer)
*/

type Animal interface {
	Eat()
}

type Cat struct {
	name string
}

type Dog string

// the implemets* its implicit just by use the same method name
// as the interface have
func (cat *Cat) Eat() {
	fmt.Println("el gato " + cat.name + " come")
}

func (dog Dog) Eat() {
	fmt.Println("el perro " + dog + " come")
}

func main() {
	cat := Cat{"PepeCat"}
	dog := Dog("AlbertDog")
	cat.Eat()
	dog.Eat()
}
