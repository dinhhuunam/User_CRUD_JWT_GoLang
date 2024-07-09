package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int       `json:"id" gorm:"primaryKey,column:id"`
	Username  string    `json:"username" gorm:"not null,unique,column:username"`
	Email     string    `json:"email" gorm:"not null,unique,column:email"`
	Password  string    `json:"password" gorm:"not null,column:password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreation struct {
	Username string `json:"username" gorm:"not null,unique,column:username"`
	Email    string `json:"email" gorm:"not null,unique,column:email"`
	Password string `json:"password" gorm:"not null,column:password"`
}
type UserRead struct {
	Id       int    `json:"id" gorm:"primaryKey,column:id"`
	Username string `json:"username" gorm:"not null,unique,column:username"`
}

func (UserRead) TableName() string {
	return "users"
}
func (UserCreation) TableName() string {
	return "users"
}
