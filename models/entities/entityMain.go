package entities

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"-" gorm:"column:password"`
	// Password  string         `json:"password" gorm:"column:password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
