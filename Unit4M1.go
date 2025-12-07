// Author: William Provost
// Version: 1.0.0
// Date: 2025-12-07
// Fileoverview: This program is a simple calculator performing multiple operations

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2, result float64
	var operation string

	fmt.Println("Welcome to my calculator program.")
	fmt.Println("Which operation would you like to perform today? (Select by typing the letter in front of the operation.)")
	fmt.Println("a. add")
	fmt.Println("b. subtract")
	fmt.Println("c. multiply")
	fmt.Println("d. divide")
	fmt.Println("e. absolute value")
	fmt.Println("f. round")
	fmt.Println("g. raise to an exponent")
	fmt.Println("h. square root")

	fmt.Print("\nWhat operation do you want to choose: ")
	fmt.Scanln(&operation)

	switch operation {
	case "a":
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)
		result = num1 + num2
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
	case "b":
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)
		result = num1 - num2
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case "c":
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)
		result = num1 * num2
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case "d":
		fmt.Print("Enter the numerator: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the denominator: ")
		fmt.Scanln(&num2)
		result = num1 / num2
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
	case "e":
		fmt.Print("Enter a number: ")
		fmt.Scanln(&num1)
		result = math.Abs(num1)
		fmt.Printf("|%.2f| = %.2f\n", num1, result)
	case "f":
		fmt.Print("Enter a number: ")
		fmt.Scanln(&num1)
		result = math.Round(num1)
		fmt.Printf("round(%.2f) = %.0f\n", num1, result)
	case "g":
		fmt.Print("Enter the base: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the exponent: ")
		fmt.Scanln(&num2)
		result = math.Pow(num1, num2)
		fmt.Printf("%.2f ^ %.2f = %.2f\n", num1, num2, result)
	case "h":
		fmt.Print("In order to calculate the square root, you will need to supply a non-negative value: ")
		fmt.Scanln(&num1)
		result = math.Sqrt(num1)
		fmt.Printf("The square root of %.0f is %.0f\n", num1, result)
	default:
		fmt.Println("Invalid operation selected.")
	}

	fmt.Println("\nDone.")
}
