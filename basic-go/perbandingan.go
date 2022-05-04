package main

import "fmt"

func main() {
	var o = "aaa"
	var p = "aaa"
	var ggg bool = o == p
	fmt.Println(ggg)

	var cc = 90
	var dd = 80
	fmt.Println(cc > dd)
	fmt.Println(cc < dd)
	fmt.Println(cc == dd)
	fmt.Println(cc != dd)
}
