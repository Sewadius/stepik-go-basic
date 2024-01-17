package main

import "fmt"

func main() {
	const Pi = 3.14
	var radius float64

	fmt.Scan(&radius)
	fmt.Println(Pi * radius * radius)
}
