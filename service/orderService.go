package service

import (
	"VueGin/model"
	"VueGin/repository"
	"VueGin/repository/query"
	"errors"
)

type OrderSrv interface {
	List(req *query.ListQuery) (Orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Order model.Order) (*model.Order, error)
	Exist(Order model.Order) *model.Order
	ExistByOrderId(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(Order model.Order) (bool, error)
	// Delete(id string) (bool, error)
}

type OrderService struct {
	Repo repository.OrderRepoInterface
}

func (srv *OrderService) List(req *query.ListQuery) (Orders []*model.Order, err error) {
	return srv.Repo.List(req)
}
func (srv *OrderService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *OrderService) Get(Order model.Order) (*model.Order, error) {
	return srv.Repo.Get(Order)
}
func (srv *OrderService) Exist(Order model.Order) *model.Order {
	return srv.Repo.Exist(Order)
}

func (srv *OrderService) ExistByOrderId(id string) *model.Order {
	return srv.Repo.ExistByOrderId(id)
}

func (srv *OrderService) Add(Order model.Order) (*model.Order, error) {
	return srv.Repo.Add(Order)
}

func (srv *OrderService) Edit(Order model.Order) (bool, error) {
	o := srv.ExistByOrderId(Order.OrderId)
	if o == nil || o.Mobile == "" {
		return false, errors.New("訂單號不存在")
	}
	return srv.Repo.Edit(Order)
}

// func (srv *OrderService) Delete(Order model.Order) (bool, error) {
// 	Order.IsDeleted = !Order.IsDeleted
// 	return srv.Repo.Delete(Order)
// }

func (srv *OrderService) Delete(Order model.Order) (bool, error) {
	Order.IsDeleted = !Order.IsDeleted
	return srv.Repo.Delete(Order)
}
