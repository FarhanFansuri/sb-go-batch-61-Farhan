package controllers

import (
	. "buku/data"
	"database/sql"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler untuk menambahkan buku baru
func createBuku(c *gin.Context) {

	var buku Buku
	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buku.CreatedAt = time.Now()
	buku.ModifiedAt = time.Now()

	if err := db.Create(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil ditambahkan", "data": buku})
}

// Handler untuk mendapatkan semua buku
func getAllBuku(c *gin.Context) {
	var buku []Buku
	if err := db.Find(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buku})
}

// Handler untuk mendapatkan buku berdasarkan ID
func getBukuByID(c *gin.Context) {
	id := c.Param("id")
	var buku Buku

	if err := db.First(&buku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buku})
}

// Handler untuk memperbarui data buku berdasarkan ID
func updateBuku(c *gin.Context) {
	id := c.Param("id")
	var buku Buku

	if err := db.First(&buku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buku.ModifiedAt = time.Now()
	if err := db.Save(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diperbarui", "data": buku})
}

// Handler untuk menghapus buku berdasarkan ID
func deleteBuku(c *gin.Context) {
	id := c.Param("id")
	var buku Buku

	if err := db.First(&buku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	if err := db.Delete(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}

func createDatabase() {
	dsn := "username:password@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS dbname")
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
}
