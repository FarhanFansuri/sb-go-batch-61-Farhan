package main

import (
	"fmt"
)

func main() {
	x := 12
	y := &x

	fmt.Println("nilai variabel x : ", x)
	fmt.Println("nilai variabel y : ", *y)

	fmt.Println("nilai pointer x : ", &x)
	fmt.Println("nilai pointer y : ", y)

	x = 2
	fmt.Println("nilai variabel x : ", x)
	fmt.Println("nilai variabel y : ", *y)
}
