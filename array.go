package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "R"
	names[1] = "J"
	names[2] = "a"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi array langsung
	nilai := [3]int{
		100, 90, 70,
	}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])

	// menghitung panjang array
	fmt.Println(len(names))
	fmt.Println(len(nilai))
}