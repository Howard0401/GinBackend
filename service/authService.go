package service

import (
	utils "VueGin/Utils"
	"VueGin/model"
	"VueGin/repository"
	"fmt"
)

type AuthRequest struct {
	AppKey     string `form:"app_key" binding:"required"`
	AppSecrect string `form:"app_secrect" binding:"required"`
}

type AuthSrv interface {
	// GetAuth(appKey, appSecret string) (model.Auth, error)
	// CheckAuth(Auth model.Auth) error
	GetToken(Auth model.User) (string, error)
}

type AuthService struct {
	Repo repository.AuthRepoInterface
}

//有種寫法是把CheckAuth放在Service層，GetAuth放在Service層內的dao，dao中有使用model層的方法Get，從資料庫內尋找
//跟拆成Service和Repository明顯不同，但哪個好？
func (srv *AuthService) GetToken(user model.User) (string, error) {

	nickName, err := srv.Repo.CheckUserAuth(user)
	if err != nil {
		return "", fmt.Errorf("CheckUserAuth error: %v", err)
	}

	token, err := utils.GenerateToken(nickName)
	if err != nil {
		return "", fmt.Errorf("GenerateToken error: %v", err)
	} else {
		return token, nil
	}

}

// func (srv *AuthService) CheckAuth(Auth model.Auth) error {
// 	auth := model.Auth{
// 		AppKey:    req.AppKey,
// 		AppSecret: req.AppSecrect,
// 	}
// 	result, err := srv.Repo.GetInfo(auth)
// 	// if auth.ID > 0 {
// 	// 	return nil
// 	// }

// 	return errors.New("auth info does not exist.")
// }

// // func (srv *AuthService) GetAuth(appKey, appSecret string) (model.Auth, error) {
// // 	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
// // 	return srv.Repo.GetInfo(auth)
// // }
