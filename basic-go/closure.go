package main

import "fmt"

func main() {
	counter := 0
	name := "aaa"
	/* Jadi fungsi increment ini dapat merubah
	   nilai variabel di atasnya itulah kenapa
	   fitur closure dikatakan dapat berinteraksi
	   dengan data sekitar
	   jika ingin deklarasi nama variabel sama
	   di function increment maka buat variabel
	   baru dengan nama yang sama namun jangan
	   menggunakan = saja harus :=, jika
	   menggunakan =  maka akan merubah nilai */

	increment := func() {
		name := "bbb"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	// name = "ccc"
	// name = "ccc" bisa merubah name := "aaa"
	// tetapi tidak dapat merubah nilai
	// variabel yang ada di fungsi increment

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
