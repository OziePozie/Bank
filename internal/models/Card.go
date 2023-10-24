package models

import "time"

type Card struct {
	Number         string    `json:"number" gorm:"primaryKey"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	CurrencyTag    string    `json:"CurrencyTag"`
	History        []History `json:"history" gorm:"foreignKey:ID"`
	IsCardActive   bool      `json:"isCardActive"`
}
