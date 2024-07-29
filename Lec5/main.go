package main

import (
	"fmt"
)

func main() {
	// // классический условный оператор
	// // if condition {
	// // 		body
	// // } else {
	// // }

	// var value int
	// fmt.Scan(&value)
	// if value%2 == 0 {
	// 	fmt.Println("The number", value, "is even")

	// } else {
	// 	fmt.Println("The number", value, "is odd")
	// }

	// // // if condition1 {

	// // // } else if {
	// // // 	...
	// // // } else {

	// // // }

	// var color string
	// fmt.Scan(&color)
	// if strings.Compare(color, "green") == 0 {
	// 	fmt.Println("Color is green")
	// } else if strings.Compare(color, "red") == 0 {
	// 	fmt.Println("Color is red")
	// } else {
	// 	fmt.Println("Unknown color")
	// }

	// // Инициализация в блоке условного оператора
	// // Блок присваивания только :=
	// // Инициализируемая переменная видна ТОЛЬКО внутри условия (if, if else, else if)

	// if num := 10; num%2 == 0 {
	// 	fmt.Println("Четное")
	// } else {
	// 	fmt.Println("Нечетное")
	// }

	// // Ущербно

	var a, b, c float64
	var D float64 = b*b - 4*a*c
	fmt.Scan(&a, &b, &c)
	if D > 0 {
		fmt.Print("два корня")
	} else if D == 0 {
		fmt.Print("один корень")
	} else {
		fmt.Print("корней нет")
	}

}
