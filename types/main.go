package main

import (
	"fmt"
)

// custom type
type Point struct {
	x, y int
}
type Person struct {
	name string
	age  int
}

// both of these functions do the same thing
// this is a method
func (m *Person) setName(newName string) {
	m.name = newName
}

// this is a regular function
func setName(m *Person, newName string) {
	m.name = newName
}

func main() {
	nocap()
	// sliced()
	// slices()
	// arrrays()
	// instructor()
	// pointing()

	m := Person{}

	setName(&m, "joe")
	m.setName("joe")
}

func nocap() {

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:2]
	// s = p[0:12:2]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

}

func f() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

}

func printSlice(s []int) {
	// `cap` for the capacity, it's the number of items the array from the first
	// element of the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliced() {
	// this is called a slice litteral, it builds an array then a slice referencing it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// inline struct
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	square := []Point{
		Point{1, 2},
		{2, 4},
		{},
		{1, 2},
	}
	fmt.Println(square)

}

func slices() {
	week := [7]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	fmt.Println(week)

	// slices are views into an array, they don't contain any value
	weekDays := week[:4]
	weekEnd := week[3:]
	fmt.Println(weekDays, weekEnd)

	// modifying an element of the slice modify the parent array
	// and all slices sharing the same indices are updated
	// note: here it's index `0` of the slice, not of the whole array
	weekEnd[0] = "Whateverday"
	fmt.Println(weekDays, weekEnd)
	fmt.Println(week)
}

func arrrays() {
	// how to declare an array
	var a [10]int // array of 10 int
	// arrays cannot be resized
	fmt.Printf("a=%v, type(a)=%T\n", a, a)

	primes := []int{2, 3, 5, 7, 11, 13} // without the length it's a slice
	// slices are dynamically sized
	fmt.Println(primes)

	primez := [6]int{2, 3, 5, 7, 11, 13}

	// another way to create slices, either from an array or another slice
	slice := primez[2:5]
	fmt.Println(slice)

	// slices of strings are strings
	cake := "chocolate cake"
	sliceOfCake := cake[0:9]
	anotherSliceOfCake := cake[10:]
	wholeCake := cake[:] // can omit the end index
	fmt.Printf("%q %q %q\n", sliceOfCake, anotherSliceOfCake, wholeCake)
}

func instructor() {
	// how to initialize custom type
	p := Point{1, 2}
	// can initialize a subset of the pointer, by using `name:`
	p2 := Point{y: 4} // x=0 implicitely
	p3 := Point{}     // x=0, y=0

	p.x = 4 // update a field

	ptr := &p
	ptr.y = 5 // can directly update struct field through pointer without dereferencing
	// we could also do
	// (*ptr).x = 1

	fmt.Printf("p=%v, p2=%v, p3=%v, ptr=%v\n", p, p2, p3, ptr)
}

func pointing() {
	var c int = 1
	// pointer
	var p *int = &c

	// pointer value is the address
	fmt.Printf("p=%v \n", p)
	// &variable => pointer to variable
	fmt.Printf("&c=%v \n", &c)
	// *p => value of the pointed variable
	fmt.Printf("*p=%v \n", *p)

	// can modify the value of the pointed variable through the pointer
	*p = 5
	fmt.Printf("c=%v \n", c)

}

func loopWithGoto() {
	i := 0
loop:
	if i < 3 {
		fmt.Println("i =", i)
		i++
		goto loop // jump back to label
	}
}
