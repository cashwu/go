
package main

import (
"fmt"
"container/list"
)


func main(){

	//x := list.New()
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for l := x.Front(); l != nil ; l= l.Next() {
		fmt.Println(l.Value)
	}
}
