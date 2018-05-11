package main

import "fmt"

var x, y, z int = 1, 2, 3
var str string = "str"
var isHas bool
var f1, f2 float64
var (
	a int
	b string
)

func main() {
	fmt.Println(str)
	str := "abc"
	fmt.Println(str)
	fmt.Println(str[0])

	c, d := true, "test !"

	fmt.Println("Hello World")
	fmt.Println(x, y, z)
	fmt.Println(c, d)

	z = 3.0
	y = 2.0

	f1, f2 := 3.0, 2.0

	fmt.Println(z / y)
	fmt.Println(f1 / f2)

	var complexVal complex64
	complexVal = 1.2 + 12i

	fmt.Println(complexVal)
	fmt.Println(real(complexVal))
	fmt.Println(imag(complexVal))

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
