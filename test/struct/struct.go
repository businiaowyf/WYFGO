package main

import (
	"fmt"
	"unsafe"
)

type Mystruct struct {
	a [2]uint32
	b uint64
}

type Meadow struct {
	location uint8
	number   uint32
	size     uint8
}

func main() {
	var m Meadow
	fmt.Println(unsafe.Sizeof(m))
	var n Mystruct
	fmt.Println(unsafe.Sizeof(n))
}
