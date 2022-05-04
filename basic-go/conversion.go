package main

import "fmt"

func main() {
	var nilai32 int32 = 3278
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	var name = "iii"
	var i byte = name[0]
	var istring string = string(i)
	fmt.Println(name)
	fmt.Println(istring)
}
