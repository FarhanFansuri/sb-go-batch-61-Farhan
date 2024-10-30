package main

import (
	"fmt"
)

func main() {
	a := 15
	b := 4

	fmt.Println("Penjumlahan:", Tambah(a, b))
	fmt.Println("Pengurangan:", Kurang(a, b))
	fmt.Println("Perkalian:", Kali(a, b))
	fmt.Println("Pembagian:", Bagi(a, b))
	fmt.Println("Modulus:", Modulus(a, b))
}
