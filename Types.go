package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//boolean
	a, b := true, false
	fmt.Println("a:", a, "b:", b)
	c := a && b //a and b
	fmt.Println("c:", c)
	d := a || b // a or b
	fmt.Println("d:", d)

	//signed int
	var a1 int = 89
	b1 := 95
	fmt.Println("value of a is", a1, "and b is", b1)
	fmt.Printf("type of a is %T, size of a is %d", a1, unsafe.Sizeof(a1))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b1, unsafe.Sizeof(b1)) //type and size of b

	// float
	a3, b3 := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a3, b3)
	sum := a3 + b3
	diff := a3 - b3
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	//Complex numbers

	// uses the bultin complex function
	// func complex(r, i FloatType) ComplexType
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	//String
	first := "Ayush"
	last := "Rai"
	name := first + " " + last
	fmt.Println("My name is", name)

	//Type conversion
	i := 10
	var j = float64(i) //this statement will not work without explicit conversion
	fmt.Println("j", j)
}
