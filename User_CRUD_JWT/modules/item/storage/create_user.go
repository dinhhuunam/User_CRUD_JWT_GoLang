package storage

import (
	"User_CRUD_JWT/modules/item/model"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *model.UserCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
