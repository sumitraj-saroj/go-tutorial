package main

import "fmt"

func main() {
	fmt.Println("This is loops class")
	days := []string{"Sunday", "monday", "tuesday", "wednesday", "thursday", "saturday", "friday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])

	}
	for i := range days {
		fmt.Println(days[i])
	}
	for index, value := range days {
		fmt.Printf("The index is %v and the value is %v\n", index, value)
	}
	randomValue := 1
	for randomValue < 10 {
        if randomValue == 2{
            goto lco
        }
		if randomValue == 5 {
			break

		}
		fmt.Println("The Random value is ", randomValue)
		randomValue++
	}
    // go to statements
    lco:
        fmt.Println("Jumping at chaicode.in")

}
