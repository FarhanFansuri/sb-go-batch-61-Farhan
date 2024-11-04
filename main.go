package main

import "fmt"

func main() {
	// 1. buat object product
	produk1 := Product{
		Name:  "Sabun",
		Price: 12500,
		Stock: 12,
	}

	var new_price float32 = 10500

	// 2. update price
	produk1.updatePrice(&new_price)

	// 3. cetak price
	fmt.Println(produk1.Price)

	// 4. update stok
	var new_stock uint = 20
	produk1.updateStock(&new_stock)

	// 5. Tampilkan nilai stock yang sekarang
	fmt.Println(produk1.Stock)
}

type student struct {
	name  string
	grade int
}

type Product struct {
	Name  string
	Price float32
	Stock uint
}

func (p *Product) updatePrice(harga *float32) {
	p.Price = *harga
}

func (p *Product) updateStock(stok *uint) {
	p.Stock = *stok
}

type class struct {
	kelas   string
	student student
}
