package main

import "fmt"

func main() {
	fmt.Println("This is Maps")
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Ts"] = "Typescript"
	languages["Py"] = "Python"
	fmt.Println("List of all languages ", languages)
    // loops in golang
    for key , value := range languages{
        fmt.Printf("For key %v,the value is %v\n", key, value)
    }
}
