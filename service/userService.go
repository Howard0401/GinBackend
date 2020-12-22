package service

import (
	utils "VueGin/Utils"
	"VueGin/model"
	"VueGin/repository"
	"VueGin/repository/query"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserSrv interface {
	List(req *query.ListQuery) (Users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(User model.User) (*model.User, error)
	Exist(User model.User) *model.User
	ExistByUserID(id string) *model.User
	Add(User model.User) (*model.User, error)
	Edit(User model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func (srv *UserService) List(req *query.ListQuery) (Users []*model.User, err error) {
	// if req.PageSize < 1 {
	// 	req.PageSize= config.PAGE_SIZE
	// }
	return srv.Repo.List(req)
}

func (srv *UserService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *UserService) Get(User model.User) (*model.User, error) {
	return srv.Repo.Get(User)
}

func (srv *UserService) Exist(User model.User) *model.User {
	return srv.Repo.Exist(User)
}

func (srv *UserService) ExistByUserID(id string) *model.User {
	return srv.Repo.ExistByUserID(id)
}

func (srv *UserService) Add(User model.User) (*model.User, error) {
	//檢查是否已存在
	chkMobile := srv.Repo.ExistByUserMobile(User.Mobile)
	chkUID := srv.Repo.ExistByUserID(User.UserId)
	if chkMobile != nil && chkUID != nil {
		return nil, errors.New("用戶已存在")
	}

	forCryptPassword := fmt.Sprintf("password%v@forCrypt", User.Password)

	User = model.User{
		UserId:     uuid.NewV4().String(),
		Mobile:     User.Mobile,
		NickName:   User.NickName,
		Address:    User.Address,
		Password:   utils.Md5(forCryptPassword),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted:  false,
		IsLocked:   false,
	}

	return srv.Repo.Add(User)
}

func (srv *UserService) Edit(User model.User) (bool, error) {
	if User.UserId == "" {
		return false, fmt.Errorf("請輸入UserId")
	}
	editUser := srv.Repo.ExistByUserID(User.UserId)
	if editUser == nil {
		return false, errors.New("請輸入正確的UserId")
	}
	editUser = &model.User{
		UserId:     User.UserId,
		NickName:   User.NickName,
		Mobile:     User.Mobile,
		Address:    User.Address,
		CreateTime: User.CreateTime,
		IsDeleted:  false,
		IsLocked:   false,
	}
	return srv.Repo.Edit(*editUser)
}

func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("傳入要刪除的id錯誤")
	}
	User := srv.Repo.ExistByUserID(id)
	if id == "" {
		return false, errors.New("查無此ID")
	}
	User.IsDeleted = !User.IsDeleted
	return srv.Repo.Delete(*User)
}
