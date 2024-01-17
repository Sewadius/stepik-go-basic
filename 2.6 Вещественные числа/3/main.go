package main

import "fmt"

func main() {
	var fahren float64

	fmt.Scan(&fahren)
	celius := (fahren - 32) * 5 / 9
	fmt.Println(celius)
}
