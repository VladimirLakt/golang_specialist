package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("Hello, my name is %s\nmy age is %d", "Vladimir", 38)
	////////////////////////////////////////////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	//Декларирование и инициализация пользовательских значений
	var height int = 170
	fmt.Println("My height is ", height)

	// В чем "полустрогость" типизации ?
	var uid = 123155 //здесь компилятор узнаёт тип по литералу
	fmt.Println("My uid", uid)
	fmt.Printf("%T\n", uid)

	//Декларирование и инциализация переменнох одного типа(множественный случай)
	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d", firstVar, secondVar)
	fmt.Println()

	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)
	fmt.Printf("Person name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUID)

	// Немного странного
	var a, b = 30, "Vladimir"
	fmt.Println(a, b)

	// Немного хорошего. Повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200

	// Короткая декларация (короткое объявление)
	count := 10
	count += 1
	fmt.Println(count)

	// Множественное присваивание через :=
	aArg, bArg := 10, 20
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 30, 40 - выдаст ошибку повторного декларирования
	// можно по другому
	bArg, cArg := 300, 400 //если есть одна новая переменная, то компилятор примет
	fmt.Println(aArg, bArg, cArg)

	//Пример
	width, length := 20.5, 30.3
	fmt.Printf("Min dimensional of rectangle is : %.3f\n", math.Min(width, length))

}
