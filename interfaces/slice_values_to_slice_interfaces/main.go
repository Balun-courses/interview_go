package main

import "fmt"

func main() {
	var interfaces []interface{}
	values := []int{100, 200, 300, 400, 500}

	interfaces = values

	/*
		interfaces = make([]interface{}, len(values))
		for idx := 0; idx < len(values); idx++ {
			interfaces[idx] = values[idx]
		}
	*/

	fmt.Println(interfaces...)
}
