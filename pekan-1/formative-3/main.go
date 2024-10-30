package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	panjangPP, errpp := strconv.Atoi(panjangPersegiPanjang)
	lebarPP, errlp := strconv.Atoi(lebarPersegiPanjang)

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	alas, erralas := strconv.Atoi(alasSegitiga)
	tinggi, errtinggi := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int
	if errlp == nil && errpp == nil {
		luasPersegiPanjang = panjangPP * lebarPP
		kelilingPersegiPanjang = 2 * panjangPP * lebarPP
		fmt.Println("Luas persegi panjang : ", luasPersegiPanjang)
		fmt.Println("Keliling persegi panjang : ", kelilingPersegiPanjang)
	}

	if erralas == nil && errtinggi == nil {
		luasSegitiga = alas * tinggi / 2
		fmt.Println("luas segitiga : ", luasSegitiga)
	}
	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	var indexJohn string
	var indexDoe string

	switch {
	case nilaiJohn >= 80:
		indexJohn = "A"
	case nilaiJohn >= 70 && nilaiJohn < 80:
		indexJohn = "B"
	case nilaiJohn >= 60 && nilaiJohn < 70:
		indexJohn = "C"
	case nilaiJohn >= 50 && nilaiJohn < 60:
		indexJohn = "D"
	case nilaiJohn < 50:
		indexJohn = "E"

	}

	switch {
	case nilaiDoe >= 80:
		indexDoe = "A"
	case nilaiDoe >= 70 && nilaiDoe < 80:
		indexDoe = "B"
	case nilaiDoe >= 60 && nilaiDoe < 70:
		indexDoe = "C"
	case nilaiDoe >= 50 && nilaiDoe < 60:
		indexDoe = "D"
	case nilaiDoe < 50:
		indexDoe = "E"

	}

	fmt.Println("Nilai John : ", indexJohn)
	fmt.Println("Nilai Doe : ", indexDoe)

	// soal 3
	var tanggal = 13
	var bulan = 4
	var tahun = 2002
	var bulanStr string
	switch bulan {
	case 1:
		bulanStr = "Januari"
	case 2:
		bulanStr = "Februari"
	case 3:
		bulanStr = "Maret"
	case 4:
		bulanStr = "April"
	case 5:
		bulanStr = "Mei"
	case 6:
		bulanStr = "Juni"
	case 7:
		bulanStr = "Juli"
	case 8:
		bulanStr = "Agustus"
	case 9:
		bulanStr = "September"
	case 10:
		bulanStr = "Oktober"
	case 11:
		bulanStr = "November"
	case 12:
		bulanStr = "Desember"
	}
	fmt.Printf("%d %s %d \n", tanggal, bulanStr, tahun)

	// soal 4

	var generasi = func(thn int) string {
		switch {
		case thn >= 1944 && thn <= 1964:
			return "Baby boomer"
		case thn >= 1965 && thn <= 1979:
			return "Generasi X"
		case thn >= 1980 && thn <= 1994:
			return "Millenials"
		case thn >= 1995 && thn <= 2015:
			return "Generasi Z"
		}
		return ""
	}(tahun)

	fmt.Println(generasi)

}
