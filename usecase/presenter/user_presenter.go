package presenter

import (
	"taveler/infrastructure/model"
)

type UserPresenter interface {
	ResponseUser(*model.User) any
	ResponseRegisterUser(*model.User) any
	ResponseLoginUser(*model.User) any
}
