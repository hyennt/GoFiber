package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Product struct {
	ID           uint `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

type Order struct {
	ID      uint    `json:"id"`
	User    User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
