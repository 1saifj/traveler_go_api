package repository

import "taveler/infrastructure/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id int) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int) error
	UserExists(email string) (bool, error)
}
