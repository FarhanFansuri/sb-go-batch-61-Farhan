package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

func main() {
	// 	soal 1
	showStatement(false)

	// soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// soal 3
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	// soal 4
	var phones = []string{}
	insertPhone(&phones)
	for i, d := range phones {
		fmt.Println(strconv.Itoa(i+1)+".", d)
		time.Sleep(1 * time.Second)
	}

	//soal 5
	var wg sync.WaitGroup
	wg.Add(1)

	// Menjalankan displayPhones dengan goroutine
	go displayPhones(&wg)

	// Menunggu semua goroutine selesai
	wg.Wait()

	// soal 6
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

func statement(s string, i int) {
	fmt.Println(s, i)
}

func showStatement(err bool) {
	defer statement("Golang Backend Development", 2021)
	if err {
		panic("Error test")
	}
	fmt.Println("Aplikasi Berjalan")
}

func EndApp() {
	message := recover()
	if message != nil {
		fmt.Println("Ada error :", message)
	}
}

func kelilingSegitigaSamaSisi(i int, b bool) (string, error) {
	defer EndApp()
	if i != 0 && b {
		return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(i) + " cm adalah " + strconv.Itoa(i*3) + " cm", nil
	} else if i != 0 && !b {
		return strconv.Itoa(i), nil
	} else if i == 0 && b {
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if i == 0 && !b {
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
	return "", nil
}

func tambahAngka(i int, a *int) {
	*a += i
}

func cetakAngka(a *int) {
	fmt.Println(*a)
}

func insertPhone(p *[]string) {
	data := []string{
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo",
	}

	for _, d := range data {
		*p = append(*p, d)
	}

	sort.Strings(*p)
}

func getMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
	}
	close(moviesChannel)
}

func displayPhones(wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}
