package main

import "fmt"

type ErrNegativeNumber float64

func (e ErrNegativeNumber) Error() string {
	return fmt.Sprintf("this is error:%g", e)
}

func sqrt(x float64) (float64, error) {
	if x > 1 {
		return x, nil
	} else {
		return 0, ErrNegativeNumber(x)
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))

}
