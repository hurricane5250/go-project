package model

import (
	"context"
	"time"
)

var UserModel = User{}

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:64"`
	Login     string    `json:"login" gorm:"size:64"`
	PinYin    string    `json:"pin_yin" gorm:"size:64"`
	Mail      string    `json:"mail" gorm:"size:64"`
	Token     string    `json:"token" gorm:"size:255"`
	Position  string    `json:"position" gorm:"size:255"`
	Post      string    `json:"post" gorm:"size:255"`
	Status    string    `json:"status" gorm:"size:255"`
	Mobile    string    `json:"mobile" gorm:"size:64"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (User) GetById(ctx context.Context, id int) (*User, error) {
	var user User
	err := NewOrm(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
