package repository

import "github.com/Icedroid/golang-clean-architecture/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
