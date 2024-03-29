package main

import (
	"fmt"
)

// Variadic function that finds a number in a list of integers
func find(num int, nums ...int) {
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Println()
}

// Variadic function that modifies a string slice
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	// Example of using the find function
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	// Example of using the change function with a string slice
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
