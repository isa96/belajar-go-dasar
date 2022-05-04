package main

import "fmt"

func oop(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "oop"
	}
}

func main() {
	var data interface{} = oop(1)
	fmt.Println(data)

}
