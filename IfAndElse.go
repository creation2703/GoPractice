package main

import (
	"fmt"
)

func main() {
	num := 10
	if num%2 == 0 { // checks if number is even
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}

	num2 := 11
	if num2%2 == 0 { // checks if number is even
		fmt.Println("The number", num2, "is even")
	} else {
		fmt.Println("The number", num2, "is odd")
	}

	num3 := 99
	if num3 <= 50 {
		fmt.Println(num3, "is less than or equal to 50")
	} else if num3 >= 51 && num3 <= 100 {
		fmt.Println(num3, "is between 51 and 100")
	} else {
		fmt.Println(num3, "is greater than 100")
	}

	if num4 := 10; num4%2 == 0 { //checks if number is even, scope of num4 is only inside this if-else block
		fmt.Println(num4, "is even")
	} else {
		fmt.Println(num4, "is odd")
	}

	//num := 10
	//if num % 2 == 0 { //checks if number is even
	//	fmt.Println("the number is even")
	//} error because it assumes if {}; and then else{} block, for go, both of them need to be in a single line
	//else {
	//	fmt.Println("the number is odd")
	//}
}
