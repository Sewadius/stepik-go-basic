package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var text string = "- лучшая книга!"
	var scanner *bufio.Scanner
	var book string

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	book = scanner.Text()

	fmt.Println(book, text)
}
