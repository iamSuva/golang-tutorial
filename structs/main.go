package main

import "fmt"

func main() {
	fmt.Println("Structwith go language")
   // no inheritance,no super no parent
//    suva:=User{
// 	"Suvadip",
// 	"suva@gmail.com",
// 	21,
//    }
//    fmt.Printf("details of user %+v\n",suva)
//    fmt.Printf("Name is %s and email is %s\n",suva.Name,suva.Email)
  blog :=Blogpost{
	author: Author{"suvadip","suvadip@gmail.com"},
	blogid: 1,
	title: "this is first blog",
  }
  fmt.Printf("blog details %+v\n",blog);
  blog.getEmail()
  fmt.Printf("blog details after setting email %+v\n",blog);
  
}
func (a *Blogpost) getEmail() {
	fmt.Println("the email is : ",a.author.email)
	a.author.email="sm@gmail.com"
}
type User struct{
	Name string
	Email string
	Age int
}
type Author struct{
	name  string
	email string
}
type Blogpost struct{
	blogid int
	author Author
	title string

}