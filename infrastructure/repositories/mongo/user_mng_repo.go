package mongo

import (
	"context"
	"errors"

	"github.com/Andrushk/goPush/entity"
)

type UserMngRepo struct {
	ctx context.Context
}

func (t *UserMngRepo) Get(userId string) (entity.User, error) {
	return entity.User{"123", nil }, nil
}

func (t *UserMngRepo) Update(user entity.User) error {
	return errors.New("not implemented")
}
