package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "orange"

	fmt.Println("Fruit List is ", fruitList)
	fmt.Println("Fruit List is ", len(fruitList))

	var vegList = [4]string{"tomato", "potato", "carrot"}

	fmt.Println("The vegList is: ", len(vegList))
	fmt.Println("The vegList is: ", (vegList))

}
