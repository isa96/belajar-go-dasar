package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHi(hasName HasName) {
	fmt.Println("Hi", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var aaa Person
	aaa.Name = "bbb"
	sayHi(aaa)

	cat := Animal{
		Name: "Hooman",
	}
	sayHi(cat)
}
