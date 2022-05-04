package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func sayHi(name string) {
	fmt.Println("Hi", name)
}

func sayBye(name string) {
	fmt.Println("Bye", name)
}
