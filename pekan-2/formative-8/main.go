package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// soal 1
	segitiga := &segitigaSamaSisi{alas: 10, tinggi: 15}
	persegi := &persegiPanjang{panjang: 20, lebar: 10}
	tabung := &tabung{jariJari: 7, tinggi: 14}
	balok := &balok{panjang: 10, lebar: 5, tinggi: 8}

	fmt.Println("Segitiga Sama Sisi:")
	hitungLuas(segitiga)
	hitungKeliling(segitiga)
	fmt.Println()

	fmt.Println("Persegi Panjang:")
	hitungLuas(persegi)
	hitungKeliling(persegi)
	fmt.Println()

	fmt.Println("Tabung:")
	hitungVolume(tabung)
	hitungPermukaan(tabung)
	fmt.Println()

	fmt.Println("Balok:")
	hitungVolume(balok)
	hitungPermukaan(balok)
	fmt.Println()

	// Soal 2
	phone1 := phone{
		name:   "Realme c12",
		brand:  "Realme",
		year:   2021,
		colors: []string{"Red", "Blue"},
	}
	phone1.showPhone()

	// soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	// gunakan seluruh variabel tersebut untuk menghasilkan output "hasil penjumlahan dari 6 + 8 + 12 + 14 = 40"

	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := angkaPertama[0] + angkaPertama[1] + angkaKedua[0] + angkaKedua[1]

	fmt.Printf("%v %d + %d + %d + %d = %d\n", prefix, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)

}

// soal 1 ============================================================================================================================
type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s *segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s *segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p *persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p *persegiPanjang) keliling() int {
	return (p.panjang + p.lebar) * 2
}

type tabung struct {
	jariJari, tinggi float64
}

func (t *tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t *tabung) luasPermukaan() float64 {
	return (2 * math.Pi * math.Pow(t.jariJari, 2)) + (2 * math.Pi * t.jariJari * t.tinggi)
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b *balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b *balok) luasPermukaan() float64 {
	return (2 * float64(b.panjang) * float64(b.lebar)) + (2 * float64(b.panjang) * float64(b.tinggi)) + (2 * float64(b.lebar) * float64(b.tinggi))
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

func hitungLuas(shape hitungBangunDatar) {
	println("Luas : ", shape.luas())
}

func hitungKeliling(shape hitungBangunDatar) {
	println("Keliling : ", shape.keliling())
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func hitungVolume(shape hitungBangunRuang) {
	fmt.Printf("Volume: %.2f\n", shape.volume())
}

func hitungPermukaan(shape hitungBangunRuang) {
	fmt.Printf("Permukaan: %.2f\n", shape.luasPermukaan())
}

// soal 2 ============================================================================================================================

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneInterface interface {
	showPhone()
}

func (p *phone) showPhone() {
	fmt.Println("Name :", p.name)
	fmt.Println("Brand :", p.brand)
	fmt.Println("Year :", p.year)
	fmt.Println("Colors :", p.colors)
}

// soal 3 ============================================================================================================================

func luasPersegi(sisi int, state bool) interface{} {

	if state == true && sisi != 0 {
		return "luas persegi dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(sisi*sisi) + " cm"
	} else if state == false && sisi != 0 {
		return sisi
	} else if state == true && sisi == 0 {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if state == true && sisi == 0 {
		return nil
	}
	return nil
}
