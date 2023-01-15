package models

type Login struct {
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type UserRes struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
	Tel        string `json:"tel"`
	Email      string `json:"email"`
}

type User struct {
	ID         uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Patronymic string
	Tel        string
	Email      string
	Password   string
}
