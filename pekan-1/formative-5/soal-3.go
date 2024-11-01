package main

import "fmt"

func buahFavorit(nama string, buah ...string) (statement string) {

	statement = "halo nama saya " + nama + " dan buah favorit saya adalah "
	for i, b := range buah {
		if i > 0 {
			statement += ", "
		}
		statement += fmt.Sprintf("%q", b) // %q menambahkan tanda kutip ke string
	}
	return
}
