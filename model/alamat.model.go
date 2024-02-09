package model

import "time"

type Alamat struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Jalan    string    `json:"jalan"`
	Negara string `json:"negara"`
	UserID    uint      `json:"user_id"`
	//AlamatRefer uint
	User 	User `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted uint      `json:"isDeleted"` 
}

type AlamatResponse struct {
	ID        uint      `json:"id"`
	Jalan    string    `json:"jalan"`
	Negara string `json:"negara"`
	UserID    uint      `json:"-"`
	User 	UserResponse `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted uint      `json:"isDeleted"` 
}