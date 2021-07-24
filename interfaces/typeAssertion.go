package main

import "fmt"

/*
we can assert if a interface contains a value of some type
like java with instance of but in this case we gonna get a
panic error if is not the same type
we can use a similar sintax like in maps where we test the value
and assing it to a variable then assing a boolean to a variable
like
value, ok := i.(int)
in this case if the value is int then the value varible will have it
and ok will be true, otherwise ok is false and the value variable
will be 0 on the type of the interface (could be false, nil, etc.)
*/

func main() {
	var i interface{} = "a string"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	in, ok := i.(int)
	fmt.Println(in, ok)

	f := i.(float64) // this is goona trhow a panic error
	fmt.Println(f)
}
