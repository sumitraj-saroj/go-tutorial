package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is files class")
	content := "This needs to go in file -- Chaicode.com"
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length is ", length)
	defer file.Close()
}
func readFile(filename string){
    databyte, err := os.ReadFile("file.txt")



}