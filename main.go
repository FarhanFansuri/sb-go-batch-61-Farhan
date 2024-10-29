package main

// Tambahkan ini untuk mengimpor paket strings
import (
	_ "fmt"
	"strings"
	_ "strings" // Tambahkan ini untuk mengimpor paket strings
)

// Tambahkan ini untuk mengimpor paket strings

// Tambahkan ini untuk mengimpor paket strings

func main() {
	// // Hello World
	// fmt.Println("Hello World")

	// // Variabel
	// var nama = "Farhan"
	// nama = "Farhan Fansuri"
	// fmt.Println(nama)

	// // variabel 2
	// var umur int16
	// umur = 12
	// fmt.Println(umur)

	// // Variabel 3
	// kampus := "UPN Veteran Yogyakarta"
	// fmt.Println(kampus)

	// // Konstanta
	// const suku = "JAWA!!"
	// fmt.Println(suku)

	// name := "Farhan"
	// age := 25
	// height := 5.9
	// isStudent := true

	// fmt.Printf("Nama: %s\n", name)                // %s untuk string
	// fmt.Printf("Umur: %d tahun\n", age)           // %d untuk integer
	// fmt.Printf("Tinggi: %.1f\n", height)          // %.1f untuk floating-point dengan 1 desimal
	// fmt.Printf("Status pelajar: %t\n", isStudent) // %t untuk boolean
	// fmt.Printf("tipe data umur : %T", age)

	// string
	// var sentence = `Halo nama saya \n farhan` // backtik membuat
	// println(sentence)

	// // casting data
	// var str = "Halo"
	// println(string(str[0]))

	// var rpt = strings.Repeat("na", 2)
	// println(rpt)
	nama := "rrrfarhan fansurirrr"
	println(string(strings.Split(nama, " ")[0]))
	println(strings.Trim(nama, "r"))
}
