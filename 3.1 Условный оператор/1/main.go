package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b > a {
		a = b
	}

	fmt.Println(a)
}
