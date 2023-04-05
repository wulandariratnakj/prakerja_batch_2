package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	// Gabungkan kedua array
	combined := append(arrayA, arrayB...)

	// Buat map kosong untuk menghapus duplikat
	unique := map[string]bool{}

	// Loop melalui elemen-elemen slice
	for _, value := range combined {
		// Tambahkan elemen ke map jika belum ada di dalamnya
		if _, ok := unique[value]; !ok {
			unique[value] = true
		}
	}

	// Buat slice baru untuk hasil akhir yang unik
	result := []string{}
	for key := range unique {
		result = append(result, key)
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
