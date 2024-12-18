package service

import (
	"context"
	. "gitlab.meizu.com/wujunfeng/go-nuclear/server/model"
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
