package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")
	var ptr *int
	fmt.Println("The value of pointer is ", ptr)
	// nil aayega output

	myNumber := 23
	var pttr = &myNumber
	fmt.Println("The value of this pointer is ", pttr)  // memory reference
	fmt.Println("The value of this pointer is ", *pttr) // actual value 23

	*pttr = *pttr * 2
	fmt.Println("The new value is : ", myNumber) // 46
}
