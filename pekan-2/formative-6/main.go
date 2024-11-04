package main

import (
	"fmt"
	"math"
)

func main() {
	// soal 1
	var luasLigkaran float64 = 216
	var kelilingLingkaran float64 = 23
	fmt.Println("Luas Lingkaran sebelumnya : ", luasLigkaran)
	fmt.Println("Keliling Lingkaran sebelumnya : ", kelilingLingkaran)
	perbaharui(&luasLigkaran, &kelilingLingkaran, 21)
	fmt.Println("Luas Lingkaran sesudahnya : ", luasLigkaran)
	fmt.Println("Keliling Lingkaran sesudahnya : ", kelilingLingkaran)

	//soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{}
	masukan(&buah)
	for i, b := range buah {
		fmt.Println(i+1, ".", b)
	}

	// soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data

	for i, film := range dataFilm {
		fmt.Printf("%d.", i+1)
		fmt.Printf("title : %s\n", film["title"])
		fmt.Printf("duration : %s\n", film["duration"])
		fmt.Printf("genre : %s\n", film["genre"])
		fmt.Printf("year : %s\n", film["year"])
		fmt.Println()
	}

}

func perbaharui(luasLigkaran *float64, kelilingLingkaran *float64, jari float64) {
	*luasLigkaran = math.Pi * math.Pow(jari, 2)
	*kelilingLingkaran = 2 * math.Pi * jari
}

func introduce(sent *string, nama, gender, pekerjaan, umur string) {
	if gender == "laki-laki" {
		*sent = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else {
		*sent = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
}

func masukan(buah *[]string) {
	*buah = append(*buah, "Jeruk")
	*buah = append(*buah, "Semangka")
	*buah = append(*buah, "Mangga")
	*buah = append(*buah, "Strawberry")
	*buah = append(*buah, "Durian")
	*buah = append(*buah, "Manggis")
	*buah = append(*buah, "Alpukat")
}

func tambahDataFilm(title string, duration string, genre string, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}
