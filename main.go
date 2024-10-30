package main

import (
	"fmt"
	"math"
)

func main() {
	a := 15
	b := 4
	number := 17
	var shop_total int

	// tugas kalkulator
	fmt.Println("Penjumlahan:", Tambah(a, b))
	fmt.Println("Pengurangan:", Kurang(a, b))
	fmt.Println("Perkalian:", Kali(a, b))
	fmt.Println("Pembagian:", Bagi(a, b))
	fmt.Println("Modulus:", Modulus(a, b))

	// tugas isEven
	fmt.Printf("is %d even number : %t", number, is_even(number))

	// tugas discount
	shop_total = 2000
	fmt.Printf("your total price is %d and you get discount %s", shop_total, discount(shop_total))

	// man min
	fmt.Println("\n nilai yang minimal : ", math.Max(1, 2))

}
