package models

type Bill struct {
	ID           int       `json:"ID" gorm:"primaryKey"`
	Number       string    `json:"number"`
	Limit        int       `json:"limit"`
	Cards        []Card    `json:"cards" gorm:"foreignKey:Number"`
	History      []History `json:"history" gorm:"foreignKey:ID"`
	IsBillActive bool      `json:"isBillActive"`
}
