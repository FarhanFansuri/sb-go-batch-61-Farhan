package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // Variabel global untuk koneksi database

func main() {
	// Membuat database jika belum ada
	CreateDatabase()

	// Konfigurasi koneksi ke database
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/buku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Migrasi tabel
	db.AutoMigrate(&Buku{}, &Kategori{}, &User{})

	// Setup router Gin
	router := gin.Default()

	// Routes CRUD Buku
	router.POST("/buku", CreateBuku)
	router.GET("/buku", GetAllBuku)
	router.GET("/buku/:id", GetBukuByID)
	router.PUT("/buku/:id", UpdateBuku)
	router.DELETE("/buku/:id", DeleteBuku)

	// Routes CRUD Kategori
	router.POST("/kategori", CreateKategori)
	router.GET("/kategori", GetAllKategori)

	// Jalankan server
	log.Println("Server berjalan di port 8080")
	router.Run(":8080")
}

// Membuat database jika belum ada
func CreateDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal terhubung ke MySQL: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS buku")
	if err != nil {
		log.Fatalf("Gagal membuat database: %v", err)
	}
}

// Handler untuk entitas Buku
func CreateBuku(c *gin.Context) {
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

func GetAllBuku(c *gin.Context) {
	var buku []Buku
	if err := db.Find(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buku})
}

func GetBukuByID(c *gin.Context) {
	id := c.Param("id")
	var buku Buku

	if err := db.First(&buku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buku})
}

func UpdateBuku(c *gin.Context) {
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

func DeleteBuku(c *gin.Context) {
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

// Handler untuk entitas Kategori
func CreateKategori(c *gin.Context) {
	var kategori Kategori
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kategori.CreatedAt = time.Now()
	kategori.ModifiedAt = time.Now()

	if err := db.Create(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan kategori"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil ditambahkan", "data": kategori})
}

func GetAllKategori(c *gin.Context) {
	var kategori []Kategori
	if err := db.Find(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kategori"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kategori})
}

// Struktur data untuk tabel Buku
type Buku struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}

// Struktur data untuk tabel Kategori
type Kategori struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

// Struktur data untuk tabel User
type User struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}
