package repository

import (
	utils "VueGin/Utils"
	"VueGin/model"
	query "VueGin/repository/query"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (user []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	ExistByUserMobile(mobile string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(user model.User) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []*model.User, err error) {
	// fmt.Println(req)
	//後端先處理好分頁，再直接返回給前端
	limit, offset := utils.Page(req.PageSize, req.Page)
	if err := repo.DB.Order("user_id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//Create
func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if err := repo.Exist(user); err != nil {
		return nil, fmt.Errorf("用戶已存在")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用戶註冊失敗")
	}
	return &user, nil
}

//Read
func (repo *UserRepository) Get(user model.User) (*model.User, error) {
	if err := repo.DB.Where("user_id = ?", user.UserId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//Update
func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Updates(map[string]interface{}{"nick_name": user.NickName, "mobile": user.Mobile, "address": user.Address, "create_time": user.CreateTime, "update_time": time.Now()}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//Delete
func (repo *UserRepository) Delete(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Update("is_deleted", user.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//Query User Count
func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var users []model.User
	// db := repo.DB
	// if req.Where != "" {
	// 	db = db.Where(req.Where)
	// }
	if err := repo.DB.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

//Query Exist //可以這樣改計次嗎?待確認
func (repo *UserRepository) Exist(user model.User) *model.User {
	var count int64
	err := repo.DB.Find(&user).Where("nick_name=?", user.NickName).Count(&count).Error
	if count > 0 && err != nil {
		return &user
	}
	return nil
}

//Query userID
func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User
	repo.DB.Where("user_id=?", id).First(&user)
	return &user
}

//可以這樣改計次嗎?待確認
func (repo *UserRepository) ExistByUserMobile(mobile string) *model.User {
	var count int64
	var user model.User
	repo.DB.Find(&user).Where("mobile= ?", mobile).Count(&count)
	if count > 0 {
		return &user
	}
	return nil
}
