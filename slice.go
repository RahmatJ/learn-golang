package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[4:7]
	fmt.Println("capacity slice:",cap(slice))
	fmt.Println("panjang slice:",len(slice))

	months[5] = "June"
	fmt.Println("slice setelah perubahan index 5:",slice)

	fmt.Println("month:",months)
	slice[0] = "May"
	fmt.Println("months setelah perubahan index 0 slice",months)

	// append
	var slice1 = months[5:]
	fmt.Println("slice 1:",slice1)
	// coba append 1 data kalau itu lebih panjang
	slice1New := append(slice1, "new data")
	fmt.Println("slice 1 setelah append:",slice1)
	fmt.Println("months setelah append",months)
	fmt.Println("slice1New:",slice1New)
	// coba ubah slice1New
	slice1New[1] = "is this new"
	fmt.Println("slice1 setelah perubahan index 1 slice1New:",slice1)
	fmt.Println("slice1New baru:",slice1New)

	var slice2 = months[5:8]
	fmt.Println("slice 2:",slice2)
	// coba tambah 1 data tapi belum lewat caps
	slice2New := append(slice2, "atarashi")
	fmt.Println("slice 2 setelah append",slice2)
	fmt.Println("months setelah append",months)
	fmt.Println("slice2New:",slice2New)

	// coba ubah slice2New
	slice2New[0] = "new"
	fmt.Println("slice2New setelah perubahan index 0:",slice2New)
	fmt.Println("slice2 setelah perubahan slice2New",slice2)
	fmt.Println("months setelah perubahan slice2New",months)

	// membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "R"
	newSlice[1] = "J"

	fmt.Println(newSlice)
	fmt.Println("len", len(newSlice))
	fmt.Println("cap", cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copied slice:", copySlice)

	// kalau panjangnya berbeda maka akan terpotong
	copySlice2 := make([]string, 1, cap(newSlice))
	copy(copySlice2, newSlice)
	fmt.Println("copied slice 2:", copySlice2)

	// perbedaan konstruksi array dan slice
	iniArray := [...]int{1,2,3}
	iniSlice := []int{1,2,3}
	fmt.Println("ini array:", iniArray)
	fmt.Println("ini slice:", iniSlice)
	// tidak terlalu berbeda jika dicetak
}