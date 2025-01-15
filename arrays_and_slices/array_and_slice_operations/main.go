package main

import (
	"fmt"
	"unsafe"
)

func accessToArrayElement1() {
	data := [3]int{1, 2, 3}

	idx := 4
	fmt.Println(data[idx]) // panic

	fmt.Println(data[4]) // compilation error
}

func accessToArrayElement2() {
	data := [3]int{1, 2, 3}

	idx := -1
	fmt.Println(data[idx]) // panic

	fmt.Println(data[-1]) // compilation error
}

func arrayLen() {
	data := [10]int{}
	fmt.Println(len(data)) // 10
}

func capArray() {
	var data [10]int
	fmt.Println(cap(data)) // 10
}

func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	// except arrays whose element types are incomparable types
	fmt.Println(first == second)
	fmt.Println(first != second)

	//	[<, <=, >, >=]  ->  compilation error
}

func emptyArray() {
	var data [10]byte
	fmt.Println(unsafe.Sizeof(data)) // 10

	//data == nil // compilation error
}

func zeroArray() {
	var data [0]int
	fmt.Println(unsafe.Sizeof(data)) // 0
}

func negativeArray() {
	var data [-1]int // compilation error
	_ = data
}

func arrayCreation() {
	length1 := 100
	var data1 [length1]int // compilation error
	_ = length1
	_ = data1

	const length2 = 100
	var data2 [length2]int
	_ = data2
}

func makeArray() {
	_ = make([10]int, 10) // compilation error
}

func appendToArray() {
	_ = append([10]int{}, 10) // compilation error
}

func accessToSliceElement1() {
	data := make([]int, 3)
	fmt.Println(data[4]) // panic
}

func accessToSliceElement2() {
	data := make([]int, 3, 6)
	fmt.Println(data[4]) // panic
}

func accessToElement3() {
	data := make([]int, 3, 6)
	_ = data[-1] // compilation error
}

func accessToNilSlice1() {
	var data []int
	_ = data[0] // panic
}

func accessToNilSlice2() {
	var data []int
	data[0] = 10 // panic
}

func appendToNilSlice() {
	var data []int
	data = append(data, 10) // ok
}

func rangeByNilSlice() {
	var data []int
	for range data { // ok
	}
}

func makeZeroSlice() {
	data := make([]int, 0)
	fmt.Println(len(data)) // 0
	fmt.Println(cap(data)) // 0
}

func makeSlice() {
	_ = make([]int, -5)    // compilation error
	_ = make([]int, 10, 5) // compilation error

	size := -5
	_ = make([]int, size) // panic

	size = 5
	_ = make([]int, size*2, size) // panic
}

func sliceMoreThanSize() {
	data := make([]int, 2, 6)

	slice1 := data[1:6] // ok
	_ = slice1
}

func sliceWithIncorrectIndeces() {
	data := make([]int, 2, 6)

	slice2 := data[1:7] // panic
	_ = slice2

	slice3 := data[2:1] // compilation error
	_ = slice3

	left := 2
	right := 1
	slice4 := data[left:right] // panic
	_ = slice4
}

func sliceWithNilSlice() {
	var data []int

	slice := data[:]  // ok
	slice = data[0:0] // ok
	slice = data[0:1] // panic
	_ = slice
}

func increaseCapacity() {
	data := make([]int, 0, 10)
	data = data[:10:100] // panic
}
