package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	changeStrings(&s1, &s2)
	fmt.Println(s1, s2)
}

func changeStrings(s1 *string, s2 *string) {
	*s1, *s2 = *s2, *s1
}
