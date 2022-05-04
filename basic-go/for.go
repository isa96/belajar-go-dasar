package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// Contoh dengan Tipe data Slice
	slice := []string{"aaa", "bbb", "ccc"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	//Contoh dengan Tipe data Map
	person := make(map[string]string)
	person["name"] = "aaa"
	person["title"] = "nnn"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
