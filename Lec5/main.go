package main

import (
	"fmt"
	"strings"
)

func main() {
	// классический условный оператор
	// if condition {
	// 		body
	// } else {
	// }

	var value int
	fmt.Scan(&value)
	if value%2 == 0 {
		fmt.Println("The number", value, "is even")

	} else {
		fmt.Println("The number", value, "is odd")
	}

	// // if condition1 {

	// // } else if {
	// // 	...
	// // } else {

	// // }

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

}
