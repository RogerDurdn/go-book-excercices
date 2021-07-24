package main

import "fmt"

func main() {
	dowitch("hello")
	dowitch(23)
	dowitch(1.1)

}

func dowitch(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("string value:%s \n", v)
	case int:
		fmt.Printf("int value:%v\n", v)
	default:
		fmt.Printf("i dunno the type of:%T\n", v)
	}
}
