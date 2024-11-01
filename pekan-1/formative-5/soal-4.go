package main

// buatlah closure function disini
var dataFilm = []map[string]string{}

func tambahDataFilm(title string, jam string, genre string, tahun string) {

	film := map[string]string{
		"genre": genre,
		"jam":   jam,
		"tahun": tahun,
		"title": title,
	}

	dataFilm = append(dataFilm, film)

}
