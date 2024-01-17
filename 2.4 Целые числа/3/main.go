package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	x_second := x * x
	x_four := x_second * x_second
	fmt.Println(x_second * x_four)
}
