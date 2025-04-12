package interfaces

import "telegram-bot/src/model"

type UserRepository interface {
	GetBydId(id int64) (model.User, error)
	Save(user model.User) error
	GetAll() ([]model.User, error)
}
