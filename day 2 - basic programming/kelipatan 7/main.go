package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	if n%7 == 0 {
		fmt.Println("bilangan kelipatan 7")
	} else {
		fmt.Println("bukan bilangan kelipatan 7")
	}
}
