package main

import (
	"fmt"
)

func main() {
	const a = 50
	fmt.Println(a)
	const (
		name    = "Ayush"
		age     = 20
		country = "India"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	//const a = 55 //allowed
	//a = 89       //reassignment not allowed

	//var a = math.Sqrt(4)   //allowed
	//const b = math.Sqrt(4) //not allowed because the value of a const can't be left hanging before compile time

	const n = "Sam" // a string constant
	var name1 = n   // uses the default type associated with the untyped constant
	fmt.Printf("type %T value %v", name1, name1)
	fmt.Println()
	var defaultName = "Sam" //allowed by default
	fmt.Println("default type", defaultName)
	type myString string            // creating a new type called as mytype when is an alias of the type string
	var customName myString = "Sam" //allowed
	//customName = defaultName        //not allowed because of go's strong typing methods
	fmt.Println("custom type", customName)

	const trueConst = true // boolean constant
	type myBool bool
	// var defaultBool = trueConst       //allowed
	// var customBool myBool = trueConst //allowed
	// defaultBool = customBool          //not allowed
	// sames rules as of string constants
	const a1 = 5
	var intVar int = a1
	var int32Var int32 = a1
	var float64Var float64 = a1
	var complex64Var complex64 = a1
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	// The default type of these kinds of constants can be thought of as being generated on the fly depending on the context.

	var a2 = 5.9 / 8
	fmt.Printf("a's type is %T and value is %v", a2, a2) // free assignments for numeric constants
}
