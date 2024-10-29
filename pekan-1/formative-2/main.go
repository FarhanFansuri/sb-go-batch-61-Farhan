package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	fmt.Println("Bootcamp Digital Skill Sanbercode Golang")

	// soal 2
	halo := "Halo Golang"
	fmt.Println(halo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Printf("%s %s %s %s \n", kataPertama, kataKedua, kataKetiga, kataKeempat)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	// Mengonversi string ke integer dan menangani error
	num1, err1 := strconv.Atoi(angkaPertama)
	num2, err2 := strconv.Atoi(angkaKedua)
	num3, err3 := strconv.Atoi(angkaKetiga)
	num4, err4 := strconv.Atoi(angkaKeempat)

	// Cek jika ada error saat konversi
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Terjadi kesalahan saat mengonversi angka.")
		return
	}

	// Menjumlahkan semua angka
	var total = num1 + num2 + num3 + num4
	fmt.Println(total)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	// Mengonversi angka ke string sebelum digabungkan
	fmt.Println(strings.Replace(kalimat, "halo", "Hi", 2) + " - " + strconv.Itoa(angka))
}
