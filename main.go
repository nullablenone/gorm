package main

import (
	"belajar_gorm/configs"
	"belajar_gorm/models"
	"fmt"
	"log"
)

func main() {
	configs.ConnectDB()
	err := configs.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Gagal membuat table:", err)
	}

	// menambah data
	// yesa := models.User{Nama: "muhamad yesa", Email: "yesa@gmail.com"}
	// configs.DB.Create(&yesa)

	// mengambil data
	var user models.User
	configs.DB.First(&user, 1)
	fmt.Printf("ID : %d , Nama : %s. Email : %s ", user.ID, user.Nama, user.Email)
}
