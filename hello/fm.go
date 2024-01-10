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
	zen()
}

func sub(x int, y int) (int, int) {
	return x + y, 3
}

type Shoe struct {
	size  int32
	brand string
	ctype string
}

type Inventory_Crocs struct {
	Shoe
	pices int
}

func addInventory(shoe Shoe, owner string) string {
	return shoe.brand
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Wheel is a field containing an anonymous struct
	Wheel struct {
		Radius   int
		Material string
	}
}

func getCarMode(carObj car) string {
	return carObj.Model
}

type rect struct {
	length  int
	breadth int
}

func (r rect) area() int {
	return r.length * r.breadth
}

// Type Casting in GO

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func zen() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
