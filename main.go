package main

import (
	"belajar_gorm/configs"
	"belajar_gorm/models"
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
	// var user models.User
	// configs.DB.First(&user, 1)
	// fmt.Printf("ID : %d , Nama : %s. Email : %s ", user.ID, user.Nama, user.Email)

	// mengambil semua data
	// var user []models.User
	// configs.DB.Find(&user)
	// for _, item := range user {
	// 	fmt.Printf("ID : %d , Nama : %s. Email : %s \n", item.ID, item.Nama, item.Email)
	// }

	//menagambil berdasarkan kondisi
	// var user models.User
	// configs.DB.Where("email = ? ", "nullablenone@gmail.com").First(&user)
	// fmt.Printf("ID : %d , Nama : %s. Email : %s ", user.ID, user.Nama, user.Email)

	//memperbarui data
	var user models.User
	configs.DB.First(&user, 2)
	configs.DB.Model(&user).Updates(models.User{Nama: "test", Email: "test@gmail.com"})
}
