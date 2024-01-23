package main

import "fmt"

func main() {
	var result string
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	fmt.Scan(&result)
	results = append(results, result)

	sum := 0
	for i := 0; i < len(results); i++ {
		if results[i] == "w" {
			sum += 3
		} else if results[i] == "d" {
			sum++
		}
	}

	fmt.Println(sum)
}
