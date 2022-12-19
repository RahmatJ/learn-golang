package main

import "fmt"

func main() {
	// constant declaration
	// harus disertai inisialisasi
	// tidak harus ditulis tipe data
	const firstName string = "data"
	const lastName = "kita"

	fmt.Println(firstName, lastName)

	// multiple constant
	const (
		dataKita = "kita" 
		dataMereka = "mereka"
	)

	fmt.Println(dataKita, dataMereka)
}