package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	doublePointer(&number)
	fmt.Println(number)
}

func doublePointer(x *int) {
	*x *= 2
}
