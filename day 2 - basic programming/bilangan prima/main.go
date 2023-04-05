package main

import (
	"fmt"
)

func bilPrima(n int) bool {
	if n < 2 {
		return false
	}

	// for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
	// 	if n%i == 0 {
	// 		return false
	// 	}
	// }

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	if bilPrima(n) {
		fmt.Printf("%d adalah bilangan prima\n", n)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", n)
	}

}
