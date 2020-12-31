package repository

import (
	page "VueGin/Utils/pageFormat"
	"VueGin/model"
	"VueGin/repository/query"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	Add(Order model.Order) (*model.Order, error)
	Get(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(Order model.Order) (bool, error)
	Exist(Order model.Order) *model.Order
	ExistByOrderId(id string) *model.Order
	GetTotal(req *query.ListQuery) (total int64, err error)
	List(req *query.ListQuery) (Orders []*model.Order, err error)
	CalculateItemsPrice(od model.OrderDetailItems) (decimal.Decimal, string, error)
}

func (repo *OrderRepository) List(req *query.ListQuery) (Orders []*model.Order, err error) {
	fmt.Printf("%v", req)
	limit, offset := page.Page(req.PageSize, req.Page)
	if err := repo.DB.Limit(limit).Offset(offset).Find(&Orders).Error; err != nil {
		return nil, err
	}
	return Orders, err
}

func (repo *OrderRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var orders []*model.Order
	if err := repo.DB.Model(&orders).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

//這邊要再考慮一下訂單條件
func (repo *OrderRepository) Get(Order model.Order) (*model.Order, error) {
	if Order.OrderId == "" {
		return nil, fmt.Errorf("請輸入正確的訂單號")
	}
	if err := repo.DB.Where("order_id = ?", Order.OrderId).Find(&Order).Error; err != nil {
		return nil, err
	}
	return &Order, nil
}

func (repo *OrderRepository) Exist(Order model.Order) *model.Order {
	if Order.OrderId == "" {
		return nil
	}
	repo.DB.Model(&Order).Where("order_id=?", Order.OrderId)
	return &Order
}

func (repo *OrderRepository) ExistByOrderId(id string) *model.Order {
	var Order model.Order
	repo.DB.Where("order_id=?", id).First(&Order)
	return &Order
}

func (repo *OrderRepository) CalculateItemsPrice(od model.OrderDetailItems) (decimal.Decimal, string, error) {
	tmpIdList := make([]string, 0)
	tmpIdList = append(tmpIdList, od.ItemId...)
	var product []*model.Product

	// SELECT * FROM product WHERE id IN (tmpIdList);
	// https://gorm.io/docs/query.html
	err := repo.DB.Model(&product).Find(&product, tmpIdList).Error
	if err != nil {
		return decimal.Decimal{}, "", err
	}
	// 算錢不用Decimal賠慘
	total := decimal.Decimal{}
	listStr := ""
	for i := 0; i < len(product); i++ {
		item := decimal.NewFromInt(product[i].SellingPrice.IntPart())
		count := decimal.NewFromInt(int64(od.ItemCount[i]))
		total = total.Add(item.Mul(count))
		listStr += fmt.Sprintf("%v %v ", product[i].ProductName, od.ItemCount[i])
		// list[product[i].CategoryName] = count
	}
	return total, listStr, nil
}

func (repo *OrderRepository) Add(Order model.Order) (*model.Order, error) {
	totalPrice, _, err := repo.CalculateItemsPrice(Order.OrderDetail)

	if err != nil {
		return &Order, fmt.Errorf("查詢商品價錢時失敗")
	}
	Order = model.Order{
		OrderId: uuid.NewV4().String(),
		// OrderDetail: model.OrderDetailItems{},
		UserId:      Order.UserId, //若要添加訂單，得先知道是哪位User下訂
		Mobile:      Order.Mobile,
		NickName:    Order.NickName,
		TotalPrice:  totalPrice,
		PayStatus:   Order.PayStatus,
		PayTime:     time.Time{},
		PayType:     Order.PayType,
		OrderStatus: Order.OrderStatus,
		ExtraInfo:   Order.ExtraInfo,
		UserAddress: Order.UserAddress,
		IsDeleted:   false,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	// err = repo.DB.Model(&Order).Create(&Order).Error
	err = repo.DB.Create(Order).Error
	if err != nil {
		return &Order, fmt.Errorf("添加訂單失敗")
	}
	return &Order, nil
}

func (repo *OrderRepository) Edit(Order model.Order) (bool, error) {
	if Order.OrderId == "" {
		return false, fmt.Errorf("請輸入傳入的id")
	}
	o := &model.Order{}
	if err := repo.DB.Model(o).Where("order_id=?", Order.OrderId).Updates(map[string]interface{}{
		"nick_name":    Order.NickName,
		"mobile":       Order.Mobile,
		"pay_status":   Order.PayStatus,
		"extra_info":   Order.ExtraInfo,
		"user_address": Order.UserAddress,
	}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *OrderRepository) Delete(Order model.Order) (bool, error) {
	err := repo.DB.Model(&Order).Where("order_id=?", Order.OrderId).Update("is_deleted", Order.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
