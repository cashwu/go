package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))

	v := h.Sum32()
	fmt.Println(v)

	ha := sha1.New()
	ha.Write([]byte("test"))
	bs := ha.Sum([]byte{})
	fmt.Println(bs)
}
