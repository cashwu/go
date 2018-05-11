package main

import "fmt"

func main() {

	var x [5]int

	x[4] = 100

	fmt.Println(x)

	fmt.Println(len(x))

	for i, value := range x {
		fmt.Println(i, value)
	}

	for _, value := range x {
		fmt.Println(value)
	}

	y := [4]int{
		1, 2, 3, 4,
	}

	fmt.Println(y)
}
