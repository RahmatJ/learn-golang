package main

import "fmt"

func main() {
	// basic math operation
	result := 10 + 10
	fmt.Println(result)

	// augmented assignments
	a := 10
	b := 10

	/*
	x = x + a
	sama nilainya dengan
	x += a
	*/
	a = a + 10
	b += 10

	fmt.Println(a)
	fmt.Println(b)
}