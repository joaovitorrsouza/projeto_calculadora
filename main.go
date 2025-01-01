package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the operation (using this format: 22*22):")
	var input string
	fmt.Scan(&input)

	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		panic("No valid operator found")
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		panic("Invalid input format")
	}

	num1, err1 := strconv.Atoi(parts[0])
	num2, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		panic("Invalid numbers")
	}

	result := getResult(num1, num2, operator)
	fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
}

func getResult(num1 int, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Operator not valid")
	}
}
