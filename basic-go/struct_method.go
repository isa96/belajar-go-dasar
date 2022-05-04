package main

import "fmt"

/* Cara penulisan tipe struct Upper to lower
   Misal 2 kata berarti: CustomerService */

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, "My name is", customer.Name)
}

func main() {
	var aaa Customer
	aaa.Name = "bbb"
	aaa.Address = "India"
	aaa.Age = 18

	aaa.sayHi("zzz")

}
