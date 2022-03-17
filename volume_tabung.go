package main

import "fmt"

func main() {

	// Volume Tabung
	const phi = 3.14
	var (
		jari   float32
		tinggi float32
		hasil  float32
	)
	fmt.Printf("Masukan Jari-Jari: ")
	fmt.Scan(&jari)
	fmt.Printf("Masukan Tinggi: ")
	fmt.Scan(&tinggi)
	hasil = phi * jari * jari * tinggi
	fmt.Printf("Total Volume: %.6f\n", hasil)

}
