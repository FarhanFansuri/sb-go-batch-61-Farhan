package main

import (
	"fmt"
	. "formative_9/utils"
)

func main() {
	// soal 1
	segitiga := &SegitigaSamaSisi{Alas: 10, Tinggi: 15}
	persegi := &PersegiPanjang{Panjang: 20, Lebar: 10}
	tabung := &Tabung{JariJari: 7, Tinggi: 14}
	balok := &Balok{Panjang: 10, Lebar: 5, Tinggi: 8}

	fmt.Println("Segitiga Sama Sisi:")
	HitungLuas(segitiga)
	HitungKeliling(segitiga)
	fmt.Println()

	fmt.Println("Persegi Panjang:")
	HitungLuas(persegi)
	HitungKeliling(persegi)
	fmt.Println()

	fmt.Println("Tabung:")
	HitungVolume(tabung)
	HitungPermukaan(tabung)
	fmt.Println()

	fmt.Println("Balok:")
	HitungVolume(balok)
	HitungPermukaan(balok)
	fmt.Println()

	// Soal 2
	phone1 := Phone{
		Name:   "Realme c12",
		Brand:  "Realme",
		Year:   2021,
		Colors: []string{"Red", "Blue"},
	}
	phone1.ShowPhone()

	// soal 3
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := angkaPertama[0] + angkaPertama[1] + angkaKedua[0] + angkaKedua[1]
	fmt.Printf("%v %d + %d + %d + %d = %d\n", prefix, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
}
