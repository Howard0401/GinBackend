package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderDetailItems struct {
	ItemId    []string `json:"itemId"` //id price
	ItemCount []int    `json:"itemCount"`
}

type Order struct {
	OrderId     string           `json:"orderId" gorm:"column:order_id"`
	UserId      string           `json:"userId" gorm:"column:user_id"`
	Mobile      string           `json:"mobile" gorm:"column:mobile"`
	NickName    string           `json:"nickName" gorm:"column:nick_name"`
	OrderDetail OrderDetailItems `json:"orderDetail" gorm:"-"`
	// gorm:"column:order_detail"
	TotalPrice  decimal.Decimal `json:"totalPrice" gorm:"column:total_price"`
	PayStatus   int             `json:"payStatus" gorm:"column:pay_status"`
	PayType     int             `json:"payType" gorm:"column:pay_type"`
	PayTime     time.Time       `json:"payTime" gorm:"column:pay_time"`
	OrderStatus int             `json:"orderStatus" gorm:"column:order_status"`
	ExtraInfo   string          `json:"extraInfo" gorm:"column:extra_info"`
	UserAddress string          `json:"userAddress" gorm:"column:user_address"`
	IsDeleted   bool            `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt    time.Time       `json:"createAt" gorm:"column:create_at"`
	UpdateAt    time.Time       `json:"updateAt" gorm:"column:update_at"`
}
