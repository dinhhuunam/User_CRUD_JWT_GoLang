package biz

import (
	"User_CRUD_JWT/modules/item/model"
	"context"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *model.UserCreation) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}
func (biz *createUserBiz) CreateUser(ctx context.Context, data *model.UserCreation) error {
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
