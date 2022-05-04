package main

import "fmt"

func main() {
	var hasil = 1 + 1
	fmt.Println(hasil)
	var a = 1
	var b = 1
	var c = a * b
	fmt.Println(c)

	//augmented assignment
	var i = 1
	i += 1 //i=i+1
	fmt.Println(i)
	var ii = 1
	ii -= 1 //ii=ii-1
	fmt.Println(ii)
	var iii = 1
	iii *= 1 //iii=iii*1
	fmt.Println(iii)
	var iiii = 1
	iiii /= 1 //iiii=iiii/1
	fmt.Println(iiii)
	var iiiii = 1
	iiiii %= 1 //iiiii=iiiii%1
	fmt.Println(iiiii)

	//unary operator
	i++
	fmt.Println(i)
}
