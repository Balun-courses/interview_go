package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(unsafe.Pointer(&i))
		go func() {
			fmt.Print(i)
		}()
	}

	time.Sleep(2 * time.Second)
}
