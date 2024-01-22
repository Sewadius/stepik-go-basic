package main

import "fmt"

func main() {
	var a, x, y int
	fmt.Scan(&a)
	x, y = calc(a)
	fmt.Println(x, y)
}

func calc(n int) (int, int) {
	return 2 * n, n * n
}
