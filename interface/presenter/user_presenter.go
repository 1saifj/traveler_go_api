package presenter

import "taveler/infrastructure/model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUser(*model.User) (*model.User, error)
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (u *userPresenter) ResponseUser(user *model.User) (*model.User, error) {
	//return some of user
	return user, nil
}
