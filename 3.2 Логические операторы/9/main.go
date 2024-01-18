package main

import (
	"fmt"
	"math"
)

func main() {
	var answer = "NO"
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

	if math.Abs(a-c) == math.Abs(b-d) {
		answer = "YES"
	}

	fmt.Println(answer)
}
