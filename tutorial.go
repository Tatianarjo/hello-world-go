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
		return
	}

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct")

	} else {
		fmt.Println("Incorrect!")
	}
}

//33:18
