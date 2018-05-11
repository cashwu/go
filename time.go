package main

import "fmt"
import "time"

func main() {

	p := fmt.Println

	now := time.Now()

	p(now)

	then := time.Date(2018,5,10,11,11,11,1111,time.UTC)
	p(then)

	diff := now.Sub(then)

	p(diff)
}
