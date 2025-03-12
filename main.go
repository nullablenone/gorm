package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Siswa struct {
	ID       uint `gorm:"primaryKey"`
	Nama     string
	Email    string `gorm:"unique"`
	Password string
}

func main() {
	dsn := "host=localhost user=postgres password=root dbname=sekolah port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}
	log.Println("Berhasil konek ke database")

	err = db.AutoMigrate(&Siswa{})
	if err != nil {
		log.Fatal("Gagal migrate tabel:", err)
	}
	log.Println("Tabel users berhasil dibuat!")

}
