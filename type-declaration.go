package main

import "fmt"

func main() {
	// type declaration
	type NoKTP string
	type Married bool

	var noKtp NoKTP = "10010101010"
	var isMarried Married = true
	
	fmt.Println(noKtp)
	fmt.Println(isMarried)
}