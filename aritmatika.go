// aritmatika.go
package main

// Fungsi penjumlahan dua angka
func Tambah(a int, b int) int {
	return a + b
}

// Fungsi pengurangan dua angka
func Kurang(a int, b int) int {
	return a - b
}

// Fungsi perkalian dua angka
func Kali(a int, b int) int {
	return a * b
}

// Fungsi pembagian dua angka
// Mengembalikan hasil dalam float untuk akurasi
func Bagi(a int, b int) float64 {
	if b == 0 {
		return 0 // Menghindari pembagian dengan nol
	}
	return float64(a) / float64(b)
}

// Fungsi modulus dua angka
func Modulus(a int, b int) int {
	return a % b
}
