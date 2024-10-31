package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, " - Berkualitas")
		} else if i%2 == 1 && i%3 == 0 {
			fmt.Println(i, " - I Love Coding")
		} else {
			fmt.Println(i, " - Santai")
		}
	}

	// soal 2
	for i := 1; i <= 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	sentence := kalimat[2:]
	hasil := "[" + strings.Join(sentence, " ") + "]"
	fmt.Println(hasil)

	// soal  4
	sayuran := []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for index, value := range sayuran {
		fmt.Println(index+1, ". ", value)
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	volume := 1
	for key, value := range satuan {
		fmt.Println(key, " = ", value)
		volume *= value
	}
	fmt.Println("Volume Balok : ", volume)
}
