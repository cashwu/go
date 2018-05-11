package main

import "fmt"

type geo interface {
	area() float64
	permeter() float64
}

type square struct {
	w, h float64
}

func (s square) area() float64 {
	return s.w * s.h
}

func (s square) permeter() float64 {
	return 2 * s.w * 2 * s.h
}

func measure(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permeter())
}

func main() {
	s := square{w: 3, h: 4}

	measure(s)
}
