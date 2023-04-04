package main

import "fmt"

func main() {
	var luas, a, b, t float32
	fmt.Print("Masukkan sisi atas: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan sisi bawah: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&t)
	luas = (a + b) * t / 2
	fmt.Println("Luas trapesium adalah", luas)
}
