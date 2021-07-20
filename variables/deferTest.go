package main

func main() {
	defer print("cokoro")
	defereders()
	print("my ")
}

func defereders() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
	println("on the deferes")
}
