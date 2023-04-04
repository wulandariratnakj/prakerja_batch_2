package main

import "fmt"

// func trapesium(a float64, b float64, t float64) float64 {
// 	luas := (a + b) * t / 2
// 	return luas
// }

func main() {
	// fmt.Printf("Hello Prakerja Batch 2 - Alterra Academy")

	// alas := 10
	// tinggi := 8
	// luas := alas * tinggi / 2

	// fmt.Println(luas)

	// age := 18

	// if age == 17 {
	// 	fmt.Println("remaja")
	// } else if age > 17 {
	// 	fmt.Println("remaja")
	// } else {
	// 	fmt.Println("bayi")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Hello World")
	// }

	// for j := 0; j < 4; j++ {
	// 	for i := 0; i <= 4; i++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// ****
	// ****
	// ****
	// ****

	// for j := 0; j < 4; j++ {
	// 	for i := 0; i <= j; i++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// *
	// **
	// ***
	// ****

	//luas trapesium
	// var a, b, t float64

	// fmt.Print("Masukkan sisi atas: ")
	// fmt.Scanln(&a)

	// fmt.Print("Masukkan sisi bawah: ")
	// fmt.Scanln(&b)

	// fmt.Print("Masukkan tinggi: ")
	// fmt.Scanln(&t)

	// luas := trapesium(a, b, t)
	// fmt.Printf("Luas Trapesium: %.2f\n", luas)

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
