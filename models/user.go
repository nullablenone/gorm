package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Nama  string `gorm:"type:varchar(50)"`
	Email string `gorm:"unique"`
}
