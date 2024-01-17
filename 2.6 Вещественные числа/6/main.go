package main

import (
	"fmt"
	"math"
)

func main() {
	var bits int

	fmt.Scan(&bits)
	kilobytes := float64(bits) / math.Pow(2, 13)
	fmt.Println(kilobytes)
}
