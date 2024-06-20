package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("This is JSON creation")
	EncodeJson()
}

type Course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"password"`
	Tags     []string `json:"tags"`
}

func EncodeJson() {
	allCourses := []Course{
		{"HTML CSS JS", 299, "mycode.com", "abn23", []string{"frontend", "client"}},
		{"ReactJS", 299, "mycode.com", "abn23", []string{"frontend", "client"}},
		{"MERN", 299, "mycode.com", "abn23", []string{"frontend", "client"}},
		{"Angular", 299, "mycode.com", "abn23", []string{"frontend", "client"}},
	}

	finalResult, err := json.MarshalIndent(allCourses,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalResult)
}
