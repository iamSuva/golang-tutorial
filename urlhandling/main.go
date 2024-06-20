package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=nextjs&payment=erehbjhfh"

func main() {
	fmt.Println("this is all about url handling")
    response,_ :=url.Parse(myurl);
	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.RawQuery)
	fmt.Println(response.Port())

	queryparams :=response.Query();
	fmt.Println(queryparams["course"]);
    //constructing url
	partsofurl :=&url.URL{
		Scheme: "https",
		Host: "lco.dev:3000",
        Path: "/learn",
        RawQuery: "course=nextjs",
	}
	anotherUrl :=partsofurl.String()
	fmt.Println(anotherUrl)
} 
