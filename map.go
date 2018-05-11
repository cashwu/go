package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["a"] = 1
	x["b"] = 1

	fmt.Println(x)

	delete(x, "b")
	fmt.Println(x)

	if name, ok := x["B"]; ok {
		fmt.Println(name, ok)
	}
	if name, ok := x["a"]; ok {
		fmt.Println(name, ok)
	}

	y := map[int]int{
		1: 2,
		2: 3,
	}

	fmt.Println(y[2])
}
