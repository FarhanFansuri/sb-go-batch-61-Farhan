package helper

import (
	"fmt"
	"math"
	"strconv"
)

// soal 1 ============================================================================================================================

// Struct dan Method untuk Bangun Datar dan Bangun Ruang
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s *SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s *SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p *PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p *PersegiPanjang) Keliling() int {
	return (p.Panjang + p.Lebar) * 2
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t *Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t *Tabung) LuasPermukaan() float64 {
	return (2 * math.Pi * math.Pow(t.JariJari, 2)) + (2 * math.Pi * t.JariJari * t.Tinggi)
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b *Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b *Balok) LuasPermukaan() float64 {
	return (2 * float64(b.Panjang) * float64(b.Lebar)) + (2 * float64(b.Panjang) * float64(b.Tinggi)) + (2 * float64(b.Lebar) * float64(b.Tinggi))
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

func HitungLuas(shape HitungBangunDatar) {
	println("Luas : ", shape.Luas())
}

func HitungKeliling(shape HitungBangunDatar) {
	println("Keliling : ", shape.Keliling())
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func HitungVolume(shape HitungBangunRuang) {
	fmt.Printf("Volume: %.2f\n", shape.Volume())
}

func HitungPermukaan(shape HitungBangunRuang) {
	fmt.Printf("Permukaan: %.2f\n", shape.LuasPermukaan())
}

// soal 2 ============================================================================================================================

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneInterface interface {
	ShowPhone()
}

func (p *Phone) ShowPhone() {
	fmt.Println("Name :", p.Name)
	fmt.Println("Brand :", p.Brand)
	fmt.Println("Year :", p.Year)
	fmt.Println("Colors :", p.Colors)
}

// soal 3 ============================================================================================================================

func LuasPersegi(sisi int, state bool) interface{} {
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
