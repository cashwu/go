package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) (int, int) {
	return x + y, x - y
}

func add2(x, y int) (a, b int) {
	a = y
	b = x
	return
}

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(add1(1, 2))
	fmt.Println(add2(1, 2))

}
