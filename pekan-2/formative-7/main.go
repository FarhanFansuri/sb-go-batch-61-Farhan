package main

import (
	"fmt"
)

func main() {

	// soal 1
	buah1 := buah{
		Nama:       "Nanas",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      9000,
	}
	fmt.Println(buah1)
	buah2 := buah{
		Nama:       "Jeruk",
		Warna:      "Oranye",
		AdaBijinya: true,
		Harga:      8000,
	}
	fmt.Println(buah2)

	buah3 := buah{
		Nama:       "Semangka",
		Warna:      "Hijau & Oranye",
		AdaBijinya: true,
		Harga:      10000,
	}
	fmt.Println(buah3)

	buah4 := buah{
		Nama:       "Pisang",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      5000,
	}
	fmt.Println(buah4)

	// soal 2
	segitiga1 := segitiga{
		alas:   12,
		tinggi: 15,
	}
	fmt.Println("Luas Segitiga 1 : ", segitiga1.LuasSegitiga())

	persegi1 := persegi{
		sisi: 20,
	}
	fmt.Println("Luas Persegi 1 : ", persegi1.LuasPersegi())

	persegiPanjang1 := persegiPanjang{
		panjang: 12,
		lebar:   10,
	}
	fmt.Println("Luas Persegi Panjang 1 : ", persegiPanjang1.LuasPersegiPanjang())

	// soal 3
	hpRusia := phone{
		name:   "Realme c12",
		brand:  "Realme",
		year:   2023,
		colors: []string{},
	}

	hpRusia.AddColor("merah")
	hpRusia.AddColor("biru")
	fmt.Println(hpRusia.name, "varian warnanya ada", hpRusia.colors)

	// soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, d := range dataFilm {
		fmt.Printf("%d.", i+1)
		fmt.Println("title :", d.title)
		fmt.Println("duration :", d.duration/60, "Jam")
		fmt.Println("genre :", d.genre)
		fmt.Println("year :", d.year)
		fmt.Println()
	}

}

type buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s *segitiga) LuasSegitiga() float64 {
	return float64(s.alas) * float64(s.tinggi) / 2.0
}

func (p *persegi) LuasPersegi() int {
	return p.sisi * p.sisi
}

func (pp *persegiPanjang) LuasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) AddColor(color string) {
	p.colors = append(p.colors, color)
}

type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, mov *[]movie) {
	new_data := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}
	*mov = append(*mov, new_data)
}
