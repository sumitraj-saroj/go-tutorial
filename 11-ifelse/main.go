package main

import "fmt"

func main() {
	fmt.Println("this is ifelse class")
	loginCount := 25
	var result string

	if loginCount < 4 {
		result = "Regular user"
	} else {
		result = "Something else"
	}

	fmt.Println(result)

	var num string
	if 9%2 == 0 {
		num = "Is Even"
	} else {
		num = "Is Odd"
	}
	fmt.Println(num)

}
