package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is switch case class")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice number is ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("You can play 1 step")
	case 2:
		fmt.Println("You can play 2 steps")
	case 3:
		fmt.Println("You can play 3 steps")
        fallthrough
	case 4:
		fmt.Println("You can play 4 steps ")
        fallthrough
	case 5:
		fmt.Println("You can play 5 steps")
	case 6:
		fmt.Println("You can play 6 steps and roll the dice once more")
	default:
		fmt.Println("What was that!")
	}
}
