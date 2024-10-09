package models


type User struct {
	Id int `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique:not null"`
	Password string `json:"password" gorm:"not null"`
}

