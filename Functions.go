package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice //simple function example with a single return value
}

//	func rectProps(length, width float64) (float64, float64) {
//		var area = length * width
//		var perimeter = (length + width) / 2
//		return area, perimeter
//	}
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value but auto returns area and perimeter as they are the only specified values
}
func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	area1, _ := rectProps(10.8, 5.6) // perimeter is discarded as _ is a write only identifier
	fmt.Printf("Area %f ", area1)
}
