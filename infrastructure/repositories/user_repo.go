package repositories

import "github.com/Andrushk/goPush/entity"

type UserRepo interface {
	Get(userId string) (entity.User, error)
	Update(user entity.User) error
}