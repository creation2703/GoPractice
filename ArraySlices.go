package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	var a [3]int //int array with length 3
	fmt.Println(a)
	var a1 [3]int //int array with length 3
	a1[0] = 12    // a1rra1y index sta1rts a1t 0
	a1[1] = 78
	a1[2] = 50
	fmt.Println(a1)
	a2 := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a2)
	a3 := [3]int{12}
	fmt.Println(a3)
	a4 := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a4)
	//a := [3]int{5, 78, 8}
	//var b [5]int
	//b = a //not possible since [3]int and [5]int are distinct types

	a5 := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a5 // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a5)
	fmt.Println("b is ", b)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	a6 := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a6); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a6[i])
	}

	a7 := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a7 { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

	a8 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a8)
	var b1 [3][2]string
	b1[0][0] = "apple"
	b1[0][1] = "samsung"
	b1[1][0] = "microsoft"
	b1[1][1] = "google"
	b1[2][0] = "AT&T"
	b1[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b1)

	a9 := [5]int{76, 77, 78, 79, 80}
	var b3 []int = a9[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b3)

	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	fruitarray1 := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice1 := fruitarray1[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice1), cap(fruitslice1)) //length of is 2 and capacity is 6
	fruitslice1 = fruitslice1[:cap(fruitslice1)]                                       //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice1), "and capacity is", cap(fruitslice1))

	i := make([]int, 5, 5)
	fmt.Println(i)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
