package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Crazy Quiz Game!")

	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!", name)
}

//33:18
