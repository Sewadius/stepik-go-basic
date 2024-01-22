package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(squareSumm(a, b))
}

func squareSumm(a, b int) int {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i * i
	}
	return sum
}
