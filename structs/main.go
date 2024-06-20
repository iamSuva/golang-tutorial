package main

import "fmt"

func main() {
	fmt.Println("Structwith go language")
   // no inheritance,no super no parent
   suva:=User{
	"Suvadip",
	"suva@gmail.com",
	21,
   }
   fmt.Println(suva)
   fmt.Printf("Name is %s and email is %s\n",suva.Name,suva.Email)

}
type User struct{
	Name string
	Email string
	Age int
}
