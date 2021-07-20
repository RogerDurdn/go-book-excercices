package main

func main() {
	a, b := 1, 3

	c := &a
	println(*c)
	c = &b
	println(*c)
}
