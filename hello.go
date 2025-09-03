package main

import (
	"fmt"
)

func main() {
	// // hello world
	// fmt.Println("Welcome to the playground!")

	// // simple comment
	// fmt.Println("A pithy quote:", quote.Opt())
	// fmt.Println("The time is", time.Now())

	// function call
	// variableTypes()
	// sendBirthDay()
	mathBug()
}

/**
* Multiline comment
**/
func variableTypes() {
	// initialize variables here, if you don't specify a default value, it default to zero values
	var smsSendingLimit int // 0
	var costPerSMS float64  // 0.00
	var hasPermission bool  // false
	var username string     // "" (empty string)

	// Sad variable declaration, a unused variable declaration is a compiler error
	// var mySkillIssues int
	// mySkillIssues = 42

	// easier variable declaration with infered type
	// mySkillIssues := 42

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func sendBirthDay() {
	messageStart := "Happy birthday you are now"
	age := 21

	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)
}

func mathBug() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}
