package main

import "fmt"

func main() {
	defer fmt.Println("Zero")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("three")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}
