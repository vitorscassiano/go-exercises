package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(listNumbers []int) int {
	var acc int = 0

	for _, number := range listNumbers {
		acc += number
	}
	return acc
}

func formatInputString(args []string) []int {
	var listNumbers []int

	for _, item := range args {
		number, _ := strconv.Atoi(item)
		listNumbers = append(listNumbers, number)
	}

	return listNumbers
}

func main() {
	argsWithoutProg := os.Args[1:]

	listNumbers := formatInputString(argsWithoutProg[1:])

	operator := argsWithoutProg[0]

	switch operator {
	case "add":
		result := add(listNumbers)
		fmt.Println(result)

	case "sub":
		fmt.Println("Develop the command sub")

	case "mult":
		fmt.Println("Develop the command mult")

	default:
		fmt.Println("This command does not exists.")
	}
}

// calculatorctl add 3 4
// calculatorctl multiply 2 3
