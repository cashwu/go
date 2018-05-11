
package main

import (
"fmt"
"hash/crc32"
"crypto/sha1"
)

func main(){
	h := crc32.NewIEEE()
	h.Write([]byte("test"))

	v := h.Sum32()
	fmt.Println(v)


	ha := sha1.New()
	ha.Write([]byte("test"))
	bs := ha.Sum([]byte{})
	fmt.Println(bs)
}

