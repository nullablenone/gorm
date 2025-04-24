package main

import (
	"belajar_gorm/configs"
	"belajar_gorm/models"
	"fmt"
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

	// membuat data

	// user := models.User{
	// 	Nama:  "muhamad yesa",
	// 	Email: "nullablenone@gmail.com",
	// 	Profile: models.Profile{
	// 		Bio: "im a golang developer, amin",
	// 	},
	// }

	// if test := db.Create(&user).Error; test != nil {
	// 	log.Fatal(test.Error())
	// }

	var user models.User

	//fetch data + relasinya
	// if test := db.Preload("Profile").First(&user, 1).Error; test != nil {
	// 	log.Fatal(test.Error())
	// }
	// fmt.Println(user.Nama)
	// fmt.Println(user.Profile.Bio)

	db.Model(&user.Profile).Update("bio", "bio di update")
	fmt.Println(user.Profile.Bio)
	log.Println("berhasil !")

}
