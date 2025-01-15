package main

import (
	"fmt"
	"unsafe"
)

//go:noinline
func allocation(index int) byte {
	var data [1 << 20]byte
	return data[index]
}

func main() {
	var array [10]int
	address1 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#1 array address:", address1)

	allocation(100)

	address2 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#2 array address:", address2)
	fmt.Println("#1 array address:", address1)
}
