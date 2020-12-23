package service

import (
	"VueGin/model"
	"VueGin/repository"
	"VueGin/repository/query"
	"errors"
)

type ProductSrv interface {
	List(req *query.ListQuery) (Product []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Product model.Product) (*model.Product, error)
	Exist(Product model.Product) *model.Product
	ExistByProductID(id string) (*model.Product, error)
	Add(Product model.Product) (*model.Product, error)
	Edit(Product model.Product) (bool, error)
	Delete(id string) (bool, error)
}

type ProductService struct {
	Repo repository.ProductRepoInterface
}

func (srv *ProductService) List(req *query.ListQuery) (Product []*model.Product, err error) {
	return srv.Repo.List(req)
}

func (srv *ProductService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *ProductService) Get(Product model.Product) (*model.Product, error) {
	return srv.Repo.Get(Product)
}

func (srv *ProductService) Exist(Product model.Product) *model.Product {
	return srv.Repo.Exist(Product)
}

func (srv *ProductService) ExistByProductID(id string) (*model.Product, error) {
	return srv.Repo.ExistByProductID(id)
}

func (srv *ProductService) Add(Product model.Product) (*model.Product, error) {
	return srv.Repo.Add(Product)
}

func (srv *ProductService) Edit(Product model.Product) (bool, error) {
	p, err := srv.ExistByProductID(Product.ProductId)
	if err == nil || p.ProductName == "" {
		return false, err
	}
	return srv.Repo.Edit(Product)
}

func (srv *ProductService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("id為空，參數錯誤")
	}
	p, err := srv.ExistByProductID(id)
	if err != nil {
		return false, err
	}
	p.IsDeleted = !p.IsDeleted
	return srv.Repo.Delete(*p)
}
