package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	fmt.Println("enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println("enter the operator(+ - * /):")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f  %s  %f = %f", num1, operator, num2, num1+num2)
	case "-":
		fmt.Printf("%f  %s  %f = %f", num1, operator, num2, num1-num2)
	case "*":
		fmt.Printf("%f  %s  %f = %f", num1, operator, num2, num1*num2)
	case "/":
		if num2 == 0.0 {
			println("division error")
		} else {
			fmt.Printf("%f  %s  %f = %f", num1, operator, num2, num1/num2)
		}
	default:
		fmt.Println("invalid operator")
	}

}
