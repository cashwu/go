package main

import "fmt"

func main() {

	msg := make(chan string)

	go func() { msg <- "ping1" }()
	go func() { msg <- "ping2" }()
	go func() { msg <- "ping3" }()
	go func() { msg <- "ping4" }()
	go func() { msg <- "ping5" }()

	m1 := <-msg
	m2 := <-msg
	m3 := <-msg
	m4 := <-msg
	m5 := <-msg

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println(m5)

}
