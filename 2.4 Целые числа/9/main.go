package main

import "fmt"

func main() {
	const Kopecks = 100
	var a, b, n int

	fmt.Scan(&a, &b, &n)
	price := (a*Kopecks + b) * n
	rubles := price / Kopecks
	fmt.Println(rubles, price%Kopecks)
}
