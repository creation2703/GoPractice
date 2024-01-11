package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	fmt.Println("My age is", age) // automatic initialization
	// output -> 0

	age = 10
	fmt.Println("Initial age", age)
	// output -> 10
	age = 50
	fmt.Println("New age", age) // assigning values to age
	// output -> 50

	var age1 int = 20
	fmt.Println("Initially declared age", age1) // declaration with initialization
	// output -> 20

	var age2 = 29 // type auto inferred
	fmt.Println("My age is", age2)

	var width, height int = 100, 50 //declaring multiple variables without inference

	fmt.Println("width is", width, "height is", height)

	var width1, height1 = 100, 50 //declaring multiple variables with inference

	fmt.Println("width is", width1, "height is", height1)

	var width2, height2 int // multiple variable declaration without initialization
	fmt.Println("width is", width2, "height is", height2)
	width2 = 100
	height2 = 50
	fmt.Println("new width is", width2, "new height is", height2)

	var (
		name    = "ayush"
		age3    = 20
		height3 int
	) // declaration of multiple variables in a single go(pun)

	fmt.Println("User details: Name ->", name, " Age ->", age3, " Height -> ", height3)

	a, b := 20, 30 // variables a and b
	println("a is", a, "b is", b)
	b, c := 40, 50 //  c is new, b declared
	println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values
	println("changed b is", b, "c is", c)

	a1, b1 := 145.1, 543.0
	c1 := math.Min(a1, b1) // runtime computation of a variables
	fmt.Println("Minimum value is", c1)
}
