package main

import "fmt"

func main() {
	fmt.Println("This is functions class")
	greet()
	adder(3, 5)
    prores:= proAdder(3,6,3,2,6,2,5,2,24,3,5,2,31,32,3,131,3,23,23,23,2)
    fmt.Println("Total addition is: ", prores)
}
func greet() {
	fmt.Println("Namaste from golang")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proAdder(value ...int) int {
	total := 0
	for _, value := range value {
		total += value
	}
	return total
}
