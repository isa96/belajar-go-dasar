package main

import "fmt"

func main() {
	var name = "bbb"
	switch name {
	case "aaa":
		fmt.Println("aaa")
	case "bbb":
		fmt.Println("bbb")
	default:
		fmt.Println("hi")
	}
}
