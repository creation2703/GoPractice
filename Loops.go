package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}

	i0 := 0
	for i0 <= 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i0)
		i0 += 2
	}

	i1 := 0
	for i1 <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i1)
		i1 += 2
	}
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//for {
	//	fmt.Println("Hello World")
	//}  infinite loop
}
