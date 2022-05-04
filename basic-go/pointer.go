package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	// Merubah 1 data field dengan pointer (&)
	// Note: Berlaku jika lokasi memory yang sama
	address2.City = "Bandung"
	// Merubah semua data field dengan pointer (*)
	// Note: Berlaku jika lokasi memory yang sama
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	/* new untuk membuat pengisian data baru untuk pointer
	   namun format pengisian data masih mengacu pada
	   struct */
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
}
