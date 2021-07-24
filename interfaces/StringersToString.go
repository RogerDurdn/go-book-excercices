package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Person::n-%v a-%v\n", p.name, p.age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	r := Person{"roger", 28}
	e := Person{"evelyn", 27}
	fmt.Println(r, e)
	p := IPAddr{1, 2, 3, 4}
	fmt.Println(p)

}
