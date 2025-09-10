package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const TWO_KBS = 2 * 1024

func main() {
	// ReadInt()
	// return
	// sum := 0
	// // `i` is only valid in the scope of the loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d + %d = %d\n", sum, i, sum+i)
	// 	sum += i
	// }
	// fmt.Printf("sum: %d\n", sum)

	// {
	// 	sum := 1
	// 	for sum < 1000 { // init & post statements are optional = while loop
	// 		sum += sum
	// 	}
	// 	fmt.Println(sum)

	// 	// infinite loop
	// 	for true {
	// 		sum += sum

	// 		if sum >= TWO_KBS {
	// 			break
	// 		}
	// 	}

	// 	// another infinite loop
	// 	for {
	// 		break
	// 	}

	// 	fmt.Println(sum)
	// }

	// fmt.Println(sqrt(47))

	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	// fmt.Println(facto(3))
	// fmt.Println(factoStr(0))

	// for v := 1.0; v <= 10; v++ {
	// 	mathSqrt := math.Sqrt(v)
	// 	custom := Sqrt(v, v/2)
	// 	fmt.Printf("Custom Sqrt(%v) = %v, math.Sqrt(%v) = %v, diff: %g\n", v, custom, v, mathSqrt, math.Abs(mathSqrt-custom))
	// }

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	// no need for `break`
	case "linuxx": // empty condition is ignored
	case "linux", "darwin": // multiple conditions to evaluate
		fmt.Println("Unix.")
	case "windows": // can only use values
		{
			fmt.Println("Windows.")
		}
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case func() time.Weekday { return today + 3 }(): // can use function evaluation for switch
		fmt.Println("In 3 days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch { // switch can have no condition,
	//  it behaves likes a chain of if conditions
	// it's equivalent to `switch true`
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// deferred()
	Boom()
}

func deferred() {
	deferredValue := "I am executed after the function returns"

	// defer the execution of the function after it returns
	defer fmt.Println(deferredValue) // the value is captured here,
	// but the function is only ran at the end

	deferredValue = "II am executed after the function returns"
	defer fmt.Println(deferredValue) // we can have multiple `defer`
	// but they are executed in reverse order

	time.Sleep(1e9) // time.Duration is in nanoseconds
	fmt.Println("hello")
}

func Boom() {
	defer fmt.Println("Boom 💥")

	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}

}

func Sqrt(x, initial float64) float64 {
	fmt.Printf("====== Sqrt(%v) ======\n", x)
	z := initial
	lastValue := z
	// range loop, [0, x)
	for i := 1; i <= 10; i++ {
		if x == 1 {
			z = 1
		} else {
			z -= (z*z - x) / (2 * z)
		}
		fmt.Printf("Iteration %d = %v\n", i, z)

		if (z*z == x) || math.Abs(lastValue-z) < 1e-10 {
			break
		}
		lastValue = z
	}

	return z
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if can have a short statement before the condition
	// that variable is block scoped, unusable outside of the `if`
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// v is also availabe in each `else` block
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Go does not support default values for arguments !!! 😭
func facto(x int) int {
	fmt.Printf("facto(%v) = %v * facto(%v - 1) \n", x, x, x)
	if x <= 0 {
		return 1
	}
	return x * facto(x-1)
}

// Go does not support default values for arguments !!! 😭
func factoStr(x int) (res string) {
	// fmt.Printf("facto(%g) = %g * facto(%g - 1) \n", x, x, x)
	res = fmt.Sprintf("facto(%d) = %d", x, x)
	for i := x - 1; i > 0; i-- {
		res += fmt.Sprintf(" * facto(%d)", i)
	}
	return res
}
