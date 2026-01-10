package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is silce class")
	var fruitList = []string{"apple", "peach", "banana", "dragonFruit"}
	fmt.Printf("This is type of %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println("The name of fruits are ", fruitList)
	fruitList = append(fruitList[2:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 35
	highScore[1] = 234
	highScore[2] = 54
	highScore[3] = 34
	//highScore[4] = 3243

	highScore = append(highScore, 8767, 867)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	// How to remove a value from slices basex on index 
	var courses = []string{"node","java","javascript","python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}
