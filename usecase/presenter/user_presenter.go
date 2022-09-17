package presenter

import "taveler/infrastructure/model"

type UserPresenter interface {
	ResponseUser(*model.User) (*model.User, error)
}
