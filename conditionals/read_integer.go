package main

import "fmt"

func ReadInt() {
	// Declare an integer variable to hold user input
	var input = 5

	// Prompt the user
	fmt.Print("Enter an integer: ")

	// Read input from the console
	_, err := fmt.Scan(&input)

	if err != nil {
		panic(err)
	}

	// Print the input back to the console
	fmt.Println("You entered:", input)
}
