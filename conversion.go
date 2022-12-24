package main

import "fmt"

func main() {
	// konversi antar tipe data

	// int
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	// ini bakal mendapatkan integer overflow, karena melebihi batas int8 (-128 s/d 127)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// string
	var name = "Rahmat"
	// mendapatkan byte
	var e = name[0]
	// konversi byte menjadi karakter string
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}