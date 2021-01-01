package model

type Banner struct {
	BannerId    string `json:"banner_id" gorm:"column:banner_id"`
	Url         string `json:"url" gorm:"column:url"`
	RedirectUrl string `json:"redirectUrl" gorm:"column:redirect_url"`
	Order       int    `json:"order" gorm:"column:order_by_idx"`
	// Order       int    `json:"order" gorm:"column:order"` 記住這個坑，欄位名用到關鍵字排序會受影響
	CreateUser string `json:"createUser" gorm:"column:create_user"`
	UpdateUser string `json:"updateUser" gorm:"column:update_user"`
}
