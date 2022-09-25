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
	FindByID(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	CreateUser(*model.User) (any, error)
	Login(*model.User) (any, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(int) error
	UserExists(string) (bool, error)
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

func (c *userInteractor) FindByEmail(email string) (*model.User, error) {
	return c.repository.FindByEmail(email)
}

func (c *userInteractor) CreateUser(user *model.User) (any, error) {
	user, err := c.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return c.presenter.ResponseRegisterUser(user), nil
}

func (c *userInteractor) Login(user *model.User) (any, error) {
	user, err := c.repository.Login(user)
	if err != nil {
		return nil, err
	}
	return c.presenter.ResponseLoginUser(user), nil
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
