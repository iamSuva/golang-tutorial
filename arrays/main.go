package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang");
	var fruits [4]string;
	fruits[0]="Apple"
	fruits[1]="Mango"
	fruits[3]="ornage"
	fmt.Println("fruit list is ",fruits);
	fmt.Println("fruit list is ",len(fruits));

	var veg=[3]string{"potato","beans","tomato"};
	fmt.Println("veg list is ",veg);

}