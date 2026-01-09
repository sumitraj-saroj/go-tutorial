package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// Standard date used for formatting
	fmt.Println(presentTime.Format("01-02-2006"))
	// to get day
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	// to get time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

    createdDate := time.Date(2016, time.September,1,7,5,0,0,time.UTC)
    fmt.Println("The time is ", createdDate)
    fmt.Println(createdDate.Format("01-2-2006 15:04:05 Monday"))

}
