package repository

import (
	"gorm.io/gorm"
	"taveler/infrastructure/model"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	Login(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int) error
	UserExists(email string) (bool, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) FindByID(id int) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Login(user *model.User) (*model.User, error) {
	err := u.db.Where("email = ?", user.Email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u *userRepository) DeleteUser(id int) error {
	err := u.db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UserExists(email string) (bool, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
