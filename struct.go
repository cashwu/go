package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"cash", 20})

	fmt.Println(person{"cash1", 11})

	x := person{"c", 1}

	fmt.Println(x)

	sp := &x
	fmt.Println(sp.age)

	sp.age = 9
	fmt.Println(sp.age)

	fmt.Println(x.age)
}
