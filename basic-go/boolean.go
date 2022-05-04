package main

import "fmt"

func main() {
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)
	//operasi boolean
	var bbb = 12
	var xxx = 11
	var ccc = bbb > 0
	var ddd = xxx > 0
	fmt.Println(ccc)
	fmt.Println(ddd)
	var nnn = ccc && ddd
	fmt.Println(nnn)
}
