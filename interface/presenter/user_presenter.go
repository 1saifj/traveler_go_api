package presenter

import (
	"taveler/infrastructure/model"
)

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUser(*model.User) any
	ResponseRegisterUser(*model.User) any
	ResponseLoginUser(*model.User) any
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (u *userPresenter) ResponseUser(user *model.User) any {
	//return some of user
	return user
}

func (u *userPresenter) ResponseRegisterUser(user *model.User) any {
	us := struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{}
	us.ID = user.ID
	us.FirstName = user.FirstName
	us.LastName = user.LastName
	us.Email = user.Email
	return us
}

func (u *userPresenter) ResponseLoginUser(user *model.User) any {
	return map[string]any{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"avatar":     user.Avatar,
		"role":       user.Role,
		"tokens":     user.Tokens,
	}

}
