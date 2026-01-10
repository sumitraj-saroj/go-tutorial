package main

import "fmt"

func main() {
	fmt.Println("This are structs in golang")
	// no inheritance in go; no super or parent
	sumit := User{"sumit", "sumit@hotmail.com", true, 23}
	fmt.Println(sumit)
	fmt.Printf("sumit details are: %+v\n", sumit)
	fmt.Printf("Name is %v and email is %v\n ", sumit.Name, sumit.Email)
    sumit.GetStruct()

}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

func (u User) GetStruct(){
    fmt.Println("Is user active ", u.Verified)

}