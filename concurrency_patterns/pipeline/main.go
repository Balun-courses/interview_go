package main

import "fmt"

func gen(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, number := range numbers {
			out <- number
		}
	}()

	return out
}

func mul(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for number := range in {
			out <- number * number
		}
	}()

	return out
}

func main() {
	ch := gen(1, 2, 3, 4, 5)
	for value := range mul(ch) {
		fmt.Println(value)
	}
}
