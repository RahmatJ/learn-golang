package main

import "fmt"

func main() {
	// kalau tidak langsung menginisialisasi, harus tulis tipe datanya
	var string_data string

	string_data = "rahmat"
	fmt.Println(string_data)

	// kalau langsung inisialisasi tak perlu tulis tipe data
	var name = "jayanto"
	fmt.Println(name)

	// untuk inisialisasi awal secara singkat bisa pakai := untuk menggantikan var
	// untuk assign variable selanjutnya cukup pakai =
	data := 30
	fmt.Println(data)

	data = 34
	fmt.Println(data)

	// deklarasi multi variable
	// bisa dipertimbangkan untuk readability code
	var (
		firstName = "firstName"
		lastName  = "lastName"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}