package main

import "fmt"

// penulisan else if HARUS di samping } setelah
// semua statement if selesai
func main() {
	var name = "bbb"
	if name == "aaa" {
		fmt.Println("aaa")
	} else if name == "bbb" {
		fmt.Println("hi, bbb")
	} else {
		fmt.Println("hello")
	}
}
