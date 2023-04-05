package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	freq := make(map[int]int)
	var unique []int

	// menghitung frekuensi kemunculan setiap angka
	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		freq[digit]++
	}

	// mencari angka yang muncul sekali
	for key, value := range freq {
		if value == 1 {
			unique = append(unique, key)
		}
	}

	return unique
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
