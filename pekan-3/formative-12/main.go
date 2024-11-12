package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
	IndeksNilai string `json:"indeks_nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

var passwordHash []byte

func init() {
	var err error
	passwordHash, err = bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
}

func basicAuth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || !checkPasswordHash(password, passwordHash) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}

func checkPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func addNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {

	if !basicAuth(w, r) {
		return
	}

	var nilai NilaiMahasiswa
	err := json.NewDecoder(r.Body).Decode(&nilai)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if nilai.Nilai > 100 {
		http.Error(w, "Nilai maksimal adalah 100", http.StatusBadRequest)
		return
	}

	switch {
	case nilai.Nilai >= 80:
		nilai.IndeksNilai = "A"
	case nilai.Nilai >= 70:
		nilai.IndeksNilai = "B"
	case nilai.Nilai >= 60:
		nilai.IndeksNilai = "C"
	case nilai.Nilai >= 50:
		nilai.IndeksNilai = "D"
	default:
		nilai.IndeksNilai = "E"
	}

	nilai.ID = uint(len(nilaiNilaiMahasiswa) + 1)

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilai)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nilai)
}

func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {

	if !basicAuth(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func main() {

	http.HandleFunc("/add_nilai", addNilaiMahasiswa)
	http.HandleFunc("/get_nilai", getNilaiMahasiswa)

	fmt.Println("Server berjalan di port 8080...")
	http.ListenAndServe(":8080", nil)

}
