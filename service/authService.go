package service

import (
	jwt "VueGin/Utils/jwt"
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

//有種寫法是把CheckAuth放在Service層，GetAuth放在與Service層平行的dao，而dao中使用model層的方法Get，從資料庫內尋找驗證是否存在
//跟拆成Service和Repository明顯不同，但哪個好？
func (srv *AuthService) GetToken(user model.User) (string, error) {

	userModel, err := srv.Repo.CheckUserAuth(user)
	if err != nil {
		return "", fmt.Errorf("CheckUserAuth error: %v", err)
	}
	//建構一個JWT struct 並放進Token
	token, err := jwt.NewJWT().GenerateToken(userModel)
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
