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
	// complex := complex128(2) // type conversion
	// fmt.Printf("Complex %v\n", complex)
	// mathBug()

	// var name string
	// var age int

	// n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%d: %s, %d\n", n, name, age)
	// typeCasting()

	// s := "Hello, 世界"

	// fmt.Println("Value:")
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%v %T\n", s[i], s[i]) // raw bytes
	// }
	// fmt.Println()

	// fmt.Println("Runes:")
	// for _, r := range s {
	// 	fmt.Printf("%c ", r) // Unicode characters
	// }

	staticTyping()
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

	// var complex complex64
	// complex = 2i // complex numbers uses `i` compared to `j` in python

	// Sad variable declaration, a unused variable declaration is a compiler error
	// var mySkillIssues int
	// mySkillIssues = 42

	// easier variable declaration with infered type
	// mySkillIssues := complex128(42)

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

func typeCasting() {
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

func staticTyping() {
	// can assign with `var` directly
	var username = "presidentSkroob"
	// can assign & type with `var` directly
	var password string = "12345"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}
