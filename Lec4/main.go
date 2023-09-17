package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Boolean
	var firstBoolean bool
	fmt.Println(firstBoolean)
	// Boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)
	// Numerics. Integers
	// int8, int16, int32, int64, int(соответсвует разрядности операционки)
	// uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b", a+b)
	// Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)
	//Узнаем сколько байт занимает переменная типа инт
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a)) // Type int size of 8 bytes

	// Эксперимент. При использовании короткого объявления - тип для целого числа - int  платформо-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b)) // Type int size of 8 bytes

	// Эксперимент 2. Используйте явное приведение типов при необходимости если уверены, что не произойдёт коллизии
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(first32 + int32(second64))
	fmt.Println(int64(first32) + second64) // лучше расширить маленький чем ужать большой

	// Эксперимент 3. Если проводятся арифметические над int и intX
	// то обязательно нужно использовать механизм приведения. Т.к. int != int64
	var third64 int64 = 16123414
	var fourthInt int = 156234
	fmt.Println(third64 + int64(fourthInt))

	// + - * / %

	// Аналогичным образом устроены uint8, uint16, uint64, uint

	// Numerics, Float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of floatFirst is %T and type floatSecond of is %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)

	// Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// Strings
	name := "Иван"
	lastname := "Ivanov"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string:", name, len(name))                   // Функция len() возвращает количество байт
	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name)) // если нужно количество символов в строке
	// Rune - руна. Это один utf-ный символ.
	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "CDE"
	fmt.Println(strings.Contains(totalString, subString))

	// run -> alias int32. Руна - это псевдоним int32
	var sampleRune rune
	fmt.Println(sampleRune) // выведет 0
	fmt.Printf("Rune as char %c\n", sampleRune)

	var anotherRune rune = 'Q'
	fmt.Println(anotherRune)
	fmt.Printf("Rune 'anothserRune' as char %c\n", anotherRune)

	var thirdRune rune = 23452
	fmt.Println(thirdRune)
	fmt.Printf("Rune 'thirdRune' as char %c\n", thirdRune)

	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "a"))       // first > second -> 1
	fmt.Println(strings.Compare("abcd", "abcvbcv")) // first < second -> -1
	fmt.Println(strings.Compare("abcd", "abcd"))    // first = second -> 0

	var aByte byte // alias uint8
	fmt.Println("Byte:", aByte)
}
