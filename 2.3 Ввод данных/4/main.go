package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s1, s2, s3 string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s1 = scanner.Text()
	scanner.Scan()
	s2 = scanner.Text()
	scanner.Scan()
	s3 = scanner.Text()

	fmt.Print(s3, "\n", s2, "\n", s1)
}
