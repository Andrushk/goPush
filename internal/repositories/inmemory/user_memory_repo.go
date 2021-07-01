package inmemory

import (
	"errors"

	"github.com/Andrushk/goPush/entity"
)

type UserMemoryRepo struct {
	users []entity.User
}

func (u *UserMemoryRepo) Init() {
	u.users = []entity.User{}
}

func (u *UserMemoryRepo) Get(userId entity.ID) (entity.User, error) {
	for i, user := range u.users {
		if user.Id == userId {
			return u.users[i], nil
		}
	}

	return entity.User{}, nil
}

func (u *UserMemoryRepo) Add(user entity.User) error {
	existUser, _ := u.Get(user.Id)
	if existUser.Id != "" {
		return errors.New("Пользователь уже существует")
	}

	u.users = append(u.users, user)
	return nil
}

func (u *UserMemoryRepo) Update(user entity.User) error {
	existUser, _ := u.Get(user.Id)
	if existUser.Id == "" {
		return errors.New("Пользователь не существует")
	}

	//старого удаляем, нового добавляем
	for i, candidate := range u.users {
		if candidate.Id == user.Id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			break
		}
	}

	u.Add(user)
	return nil
}

func (u *UserMemoryRepo) Count() int {
	return len(u.users)
}
