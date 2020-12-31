package repository

import (
	md5 "VueGin/Utils/crypto"
	"VueGin/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

type AuthRepoInterface interface {
	CheckUserAuth(User model.User) (model.User, error)
}

func (repo *AuthRepository) CheckUserAuth(User model.User) (model.User, error) {
	var u model.User
	forCryptPassword := fmt.Sprintf("password%v@forCrypt", User.Password)
	err := repo.DB.Where("nick_name=? AND password=?", User.NickName, md5.Md5(forCryptPassword)).First(&u).Error
	// err := repo.DB.Where("nick_name=? AND password=?", User.NickName, forCryptPassword).First(&u).Error
	if err != nil {
		return model.User{}, errors.New("帳號密碼錯誤")
	}
	return u, nil
}
