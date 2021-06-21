package repositories

import "github.com/Andrushk/goPush/entity"

type UserRepo interface {
	Get(userId entity.ID) (entity.User, error)
	Add(user entity.User) error
	Update(user entity.User) error
}