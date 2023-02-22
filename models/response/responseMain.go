package response

import "time"

type MessageMain struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"-" gorm:"column:password"`
	// Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"index,column:beast_id"`
}
