package main

import "fmt"

func divide(lhs, rhs int) int {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	return lhs / rhs
}

func main() {
	_ = divide(1000, 0)
}
