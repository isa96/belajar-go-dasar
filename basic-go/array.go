package main

import "fmt"

func main() {
	//manual
	var names [3]string
	names[0] = "aaa"
	names[1] = "bbb"
	names[2] = "ccc"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//langsung
	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	//Panjang Array
	fmt.Println(len(values))
	fmt.Println(len(names))
}
