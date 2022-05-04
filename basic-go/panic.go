package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi jalan")
}

func main() {
	runApp(true)

}
