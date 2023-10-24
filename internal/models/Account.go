package models

type Account struct {
	ID         int    `json:"ID" gorm:"primaryKey"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Login      string `json:"login" gorm:"unique"`
	Password   string `json:"password"`
	Bill       []Bill `json:"bills" gorm:"foreignKey:ID"`
}

type AccountDetails struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}
