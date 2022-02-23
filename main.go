package main

import (
	"errors"
	"fmt"
)

type Amount struct {
	value int
}

func NewAmount(value int) (Amount, error) {
	if value < 0 {
		return Amount{}, errors.New("Invalid amount")
	}

	return Amount{value: value}, nil
}

func (a Amount) GetValue() int {
	return a.value
}

func main() {
	amount, err := NewAmount(10)

	if err != nil {
		fmt.Println("There is no amount setted. err = ", err.Error())
	}
	fmt.Println(amount.GetValue())
}
