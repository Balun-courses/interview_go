package main

func main() {
	var x = []string{"A", "M", "C"}

	for i, s := range x {
		println(i, s)

		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
