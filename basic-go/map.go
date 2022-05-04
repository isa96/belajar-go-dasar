package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "aaa",
		"address": "bbb",
	}

	person["title"] = "ccc"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "aaa"
	book["author"] = "bbb"
	book["ooo"] = "lll"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ooo")
	fmt.Println(book)
	fmt.Println(len(book))
}
