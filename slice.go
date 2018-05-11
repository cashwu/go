package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	s2 := append(s1, 4, 5)
	fmt.Println(s2)

	s3 := make([]int, 2)
	copy(s3, s1)
	fmt.Println(s3)
}
