package presenter

import "github.com/Icedroid/golang-clean-architecture/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
