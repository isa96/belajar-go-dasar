package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
	}

	var slice1 = months[2:3]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[4:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "aaa")
	fmt.Println(slice3)
	slice3[1] = "bbb"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	//Buat Slice
	// := pengganti tulisan var tujuannya agar tahu bahwa nama tersebut adalah sebuah variabel
	newSlice := make([]string, 2, 5)
	newSlice[0] = "aaa"
	newSlice[1] = "bbb"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
