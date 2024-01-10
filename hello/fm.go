package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// s := fmt.Sprintf("I am x year ago %v ", 10)
	x, _ := sub(1, 2)
	fmt.Println(x)
	fmt.Println(sub(1, 4))
	zen()
	throwErros()
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

// Interfaces are not classes, they are slimmer.
// Interfaces don’t have constructors or deconstructors that require that data is created or destroyed.
// Interfaces aren’t hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.
// Interfaces define function signatures, but not underlying behavior. Making an interface often won’t DRY up your code in regards to struct methods. For example,
// if five types satisfy the fmt.Stringer interface, they all need their own version of the String() function.

//ERRORS IN GO

func throwErros() {
	i, err := strconv.Atoi("42b")
	for i := 0; i < 10; i++ {
		fmt.Fprintln(os.Stdout, []any{"heyy %c", i}...)
	}
	var zerr error = errors.New("something went wrong")
	fmt.Println(zerr)
	if err != nil {
		fmt.Println("couldn't convert:", err)
		fmt.Println(i)
		return
	}
}
