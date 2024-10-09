package models


type Books struct{
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	AuthorName string `json:"author_name"`
}