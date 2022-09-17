package interactor

import (
	"taveler/infrastructure/model"
	"taveler/usecase/presenter"
	"taveler/usecase/repository"
)

type userInteractor struct {
	repository repository.UserRepository
	presenter  presenter.UserPresenter
}

type UserInteractor interface {
	FindAll() ([]*model.User, error)
	FindByID(id int) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int) error
	UserExists(email string) (bool, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{
		repository: r,
		presenter:  p,
	}
}

func (c *userInteractor) FindAll() ([]*model.User, error) {
	return c.repository.FindAll()
}

func (c *userInteractor) FindByID(id int) (*model.User, error) {
	return c.repository.FindByID(id)
}

func (c *userInteractor) CreateUser(user *model.User) (*model.User, error) {
	user, err := c.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return c.presenter.ResponseUser(user)
}

func (c *userInteractor) UpdateUser(user *model.User) (*model.User, error) {
	return c.repository.UpdateUser(user)
}

func (c *userInteractor) DeleteUser(id int) error {
	return c.repository.DeleteUser(id)
}

func (c *userInteractor) UserExists(email string) (bool, error) {
	return c.repository.UserExists(email)
}
