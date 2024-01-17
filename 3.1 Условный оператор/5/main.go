package main

import "fmt"

func main() {
	var a, b float64
	var result string = "NO"
	fmt.Scan(&a, &b)

	check := int(a) / int(b)

	if float64(check) == a/b {
		result = "YES"
	}

	fmt.Println(result)
}
