package main

import (
	"belajar_gorm/configs"
	"belajar_gorm/models"
	"log"
)

func main() {
	configs.ConnectDB()
	err := configs.DB.AutoMigrate(&models.User{})
	if err != nil{
		log.Fatal("Gagal membuat table:", err)
	}
}
