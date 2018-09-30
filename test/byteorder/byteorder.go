package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	a := 0x11223344
	fmt.Printf("%x\n", a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(unsafe.Sizeof(a))

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(a))
	fmt.Println("buf=", hex.EncodeToString(buf))

	out := binary.LittleEndian.Uint32(buf)
	fmt.Printf("%x\n", out)
}
