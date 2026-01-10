package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://google.com:300"

func main() {
	fmt.Println("Welcome to url handling")
    fmt.Println(myurl)
    //parsing
    result, err :=url.Parse(myurl)
    if err != nil{
        panic(err)
    }
    fmt.Println(result.Scheme)
    fmt.Println(result.Host)
    fmt.Println(result.Path)
    fmt.Println(result.RawQuery)
    fmt.Println(result.Port())

    qparam := result.Query()
    fmt.Printf("The type of query params are: %T\n",qparam)

    // pass the reference 
    partsOfUrl := &url.URL{
        Scheme: "https",
        Host: "lco.dev",
        Path: "/tutcss",
        RawPath: "user=sumit",

    }
    anotherURL := partsOfUrl.String()
    fmt.Println(anotherURL)

}
