package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	myPrint(x, y)
}

func myPrint(val ...int) {
	for _, v := range val {
		fmt.Println(v)
	}
}
