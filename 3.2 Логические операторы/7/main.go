package main

import "fmt"

func main() {
	var age int
	var answer string
	fmt.Scan(&age)

	if age <= 13 {
		answer = "детство"
	} else if age > 13 && age <= 24 {
		answer = "молодость"
	} else if age > 24 && age <= 59 {
		answer = "зрелость"
	} else {
		answer = "старость"
	}

	fmt.Println(answer)
}
