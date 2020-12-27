package model

type Auth struct {
	*User
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}
