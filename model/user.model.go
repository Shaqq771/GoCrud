package model

import "time"

type User struct {
	ID        uint 		`gorm:"primaryKey" json:"id"`
	Name      string 	`json:"name"`
	Email 	  string `json:"email"`
	// AlamatRefer uint
	// Alamats 	Alamat `gorm:"foreignKey:AlamatRefer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted uint      `json:"isDeleted"` 
}

type UserResponse struct {
	ID        uint 		`gorm:"primaryKey" json:"id"`
	Name      string 	`json:"name"`
	Email 	  string `json:"email"`
	IsDeleted uint      `json:"isDeleted"` 
}