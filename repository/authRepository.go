package repository

import (
	"VueGin/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

type AuthRepoInterface interface {
	CheckUserAuth(User model.User) (string, error)
}

func (repo *AuthRepository) CheckUserAuth(User model.User) (string, error) {
	var u model.User
	forCryptPassword := fmt.Sprintf("password%v@forCrypt", User.Password)
	err := repo.DB.Model(&u).Where("nick_name=? AND passward=?", User.NickName, forCryptPassword).Error
	if err != nil {
		return "", errors.New("帳號密碼錯誤")
	}
	return u.NickName, nil
}
