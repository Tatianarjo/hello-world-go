package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Crazy Quiz Game!")

	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
	}

}

//33:18
