package main

import (
	"errors"
	"fmt"
)

func main() {
	// s := fmt.Sprintf("I am x year ago %v ", 10)
	x, _ := sub(1, 2)
	fmt.Println(x)
	fmt.Println(sub(1, 4))
}

func sub(x int, y int) (int, int) {
	return x + y, 3
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

type car struct {
	Make string
	Model string
	Height int
	Width int
	// Wheel is a field containing an anonymous struct
	Wheel struct {
	  Radius int
	  Material string
	}
  }
  

func getCarMode(carObj car) string {
	return carObj.Model