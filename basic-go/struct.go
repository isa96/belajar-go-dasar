package main

import "fmt"

/* Cara penulisan tipe struct Upper to lower
   Misal 2 kata berarti: CustomerService */

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var aaa Customer
	aaa.Name = "bbb"
	aaa.Address = "India"
	aaa.Age = 18

	fmt.Println(aaa.Name)
	fmt.Println(aaa.Address)
	fmt.Println(aaa.Age)

	// Opsi penulisan struct
	ccc := Customer{
		Name:    "ddd",
		Address: "Indonesia",
		Age:     15,
	}
	fmt.Println(ccc)

	// Opsi ini mudah error karena harus urut fieldnya
	eee := Customer{"fff", "Papua", 40}
	fmt.Println(eee)
}
