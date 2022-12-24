package main

import "fmt"

func main() {

	name1 := "RJ"
	name2 := "As"

	stringEqual := name1 == name2
	stringMoreThan := name1 > name2

	fmt.Println(stringEqual)
	fmt.Println(stringMoreThan)
}