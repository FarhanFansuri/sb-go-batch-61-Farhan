package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// 	soal 1
	showStatement(false)

	// soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

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
