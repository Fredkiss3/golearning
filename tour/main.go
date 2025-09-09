// Every go program is made up of packages and the `main` is the entrypoint to
// run an executable
package main

/**
A note about why a basic Go "hello world" weighs 2.3MB
Every go program include the go runtime (the garbage collector), and
other go utilities that allow for handling stack traces and panic.

we can build go apps with flags `go build -ldflags="-s -w"` which disables
debug symbols (allowing for pointing an error to an exact line number)

more info:
> http://golang.org/doc/faq#Why_is_my_trivial_program_such_a_large_binary
> https://stackoverflow.com/questions/28576173/reason-for-huge-size-of-compiled-executable-of-go
**/

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"

	// "crypto/rand" // cannot import 2 packages ending in the same name
	// because the last argument of the package is the name of the package
	// all files in `math/rand` start with `package rand`
	// but you can give it an alias like this
	cryptorand "crypto/rand" // used like `cryptorand.<function>`
)

// you can have multiple imports but it's not recommended
// import "fmt"
// import "math"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// `%g` choose the best formatting for floating points with the best decimal precision
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	a := 123.456
	b := 0.00001234

	fmt.Printf("%g\n", a) // → 123.456
	fmt.Printf("%g\n", b) // → 1.234e-05

	// Using crypto/rand
	c := make([]byte, 4)
	cryptorand.Read(c)
	fmt.Println("crypto/rand bytes:", c)

	// Every function/variable/type/struct that starts with a capital letter is exported
	// names with with `lowercase` are not exported at all (ex: `math.pi` is private)
	fmt.Printf("The value of pi is %g\n", math.Pi)

	// can define multiple variables at once
	x, y := "hello", "world"
	x, y = swap(x, y)
	fmt.Println(x, y)

	var h uint16 = 100
	for i := 0; i < 101; i++ {
		h -= 1
	}

	var k int16 = 1
	for i := 0; i < 65536; i++ {
		k -= 1
	}
	fmt.Printf("h: %v k: %v\n", h, k)

	w, t := split(10)
	fmt.Println("Named return:", w, t)

	// can declare multiple variables of the same type in one line
	// var i, j int = 1, 2

	// can declare multiple variables of different types, if we don't specify the type
	var cc, python, java = true, false, "no!"
	fmt.Println(cc, python, java)

	// ca use blocks, and within blocks, variables can be redefined
	{
		fmt.Printf("Type: %T Value: %v\n", w, w)

		// can declare multiple variables in one go, in multiple lines
		var (
			ToBe   bool       = false
			MaxInt uint64     = 1<<64 - 1
			w      complex128 = cmplx.Sqrt(-5 + 12i)
		)

		fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("Type: %T Value: %v\n", w, w)
	}

	{
		// no valye = `zero` value
		var i int     // 0
		var f float64 // 0
		var b bool    // false
		var s string  // ""
		var t rune    // \x00
		// %q = string with quotes ?
		fmt.Printf("%d %f %v %q %q\n", i, f, b, s, t)
	}

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))
}

// how to define a function,
func add(f func(int, int) int, x int) func() int {
	// return x + y
	return func() int { return f(1, 2) + x }
}

// when two or more argument follow each other,
// we can omit the types of the all args (except the last)
func sum(a, b, c int) int {
	return a + b + c
}

// functions can return any number of results
// but they need to be destructured
func swap(x, y string) (string, string) {
	return y, x
}

// go functions can have named return values,
// they appear as variables usable in the function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // x, y returned implicitely
	// naked returns are not recommended in longer functions
}
