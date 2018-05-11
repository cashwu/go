package main

import "fmt"

func choose(i int) {

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("@")

	}

}

func main() {

	choose(1)
	choose(3)

}
