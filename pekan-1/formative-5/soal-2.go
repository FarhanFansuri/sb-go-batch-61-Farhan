package main

func introduce(nama, gender, pekerjaan, umur string) (statement string) {
	if gender == "laki-laki" {
		statement = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else {
		statement = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
	return
}
