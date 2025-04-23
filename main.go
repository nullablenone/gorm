package main

import (
	"belajar_gorm/configs"
	"belajar_gorm/models"
	"log"
)

func main() {
	db := configs.ConnectDB()
	err := configs.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Gagal membuat table:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		log.Fatal("gagal membuat relasi")
	}

	user := models.User{
		Nama:  "muhamad yesa",
		Email: "nullablenone@gmail.com",
		Profile: models.Profile{
			Bio: "im a golang developer, amin",
		},
	}

	if test := db.Create(&user).Error; test != nil {
		log.Fatal(test.Error())
	}

	log.Println("berhasil !")

}
