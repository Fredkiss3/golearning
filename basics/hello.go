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

	// constants()
	// computedConstants()
	stringFormatting()
}

func stringFormatting() {
	// %v means format as the Go-style value
	s := fmt.Sprintf("I am %v years old", 10)
	// I am 10 years old
	fmt.Println(s)

	s = fmt.Sprintf("I am %v years old", "way too many")
	// I am way too many years old
	fmt.Println(s)

	// %s means the formatted string expects a string
	fmt.Printf("hello %s \n", 1)
	// will print `hello %!s(int=1)` meaning a formatting error happened
	// `%!s` => bad `%s` usage
	// `(int=1)` => You passed me an `int` instead of `string`

	// %d for integers
	fmt.Printf("hello %d \n", 1)

	// %f for floats
	PI := 3.141592653589793
	fmt.Printf("hello %f \n", PI)
	// can pass a number after % for decimals (`%.xf`)
	// it will round the number to the number of the decimals specified
	fmt.Printf("hello %.3f \n", PI)

	const name = "Saul Goodman"
	const openRate = 30.5

	// ?
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
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

func sameLineDeclaration() {
	// You can declare multiple variables on the same line:
	averageOpenRate, displayMessage := .25, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)

}

func constants() {
	// this is how we assign constants
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// we cannot assign the constant with a walrus (`:=`) operator
	// const walrus := 1

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}

func computedConstants() {
	const secondsInMinute = 60
	const minutesInHour = 60
	// constants can be computed,
	// but only if the computation can be done at compile time
	// so effectively constants can only be computed from other constants
	const secondsInHour = secondsInMinute * minutesInHour

	// `value`` is subject to change, cannot be computed at compile time
	// value := 1
	// const cc = value*1

	// the current time can only be known when the program is running
	// const currentTime = time.Now()

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)

}
