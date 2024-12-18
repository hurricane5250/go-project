package service

import (
	"context"
	. "github.com/hurricane5250/MyGoProject/server/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (r UserService) GetById(ctx context.Context, id int) (*User, error) {
	user, err := UserModel.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
