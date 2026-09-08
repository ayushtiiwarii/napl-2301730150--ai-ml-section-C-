package main

import (
	"fmt"

	"lab2/mathutil"
	"lab2/strop"
)

func main() {

	var text string
	var number int
	var b int
	var p int

	// String input
	fmt.Print("Enter a string: ")
	fmt.Scan(&text)

	fmt.Println("Reversed String:", strop.Reverse(text))
	fmt.Println("Number of Vowels:", strop.CountVowels(text))

	// Factorial input
	fmt.Print("Enter a number for factorial: ")
	fmt.Scan(&number)

	fmt.Println("Factorial:", mathutil.Factorial(number))

	// Power input
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Print("Enter p: ")
	fmt.Scan(&p)

	fmt.Println("Power:", mathutil.Power(b, p))
}
