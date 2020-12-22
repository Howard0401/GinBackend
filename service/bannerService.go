package service

import (
	"VueGin/model"
	"VueGin/repository"
	"VueGin/repository/query"
)

//基本上，service層是可以用於調用各個Repository撰寫客製化邏輯，但這邊因為只做展示，故僅套用本身的Repository服務
type BannerSrv interface {
	Add(Banner model.Banner) (*model.Banner, error)
	Get(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(id string) (bool, error)

	List(req *query.ListQuery) (Banners []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Exist(Banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
}

type BannerService struct {
	Repo repository.BannerRepoInterface
}

func (srv *BannerService) Add(Banner model.Banner) (*model.Banner, error) {
	return srv.Repo.Add(Banner)
}

func (srv *BannerService) Get(Banner model.Banner) (*model.Banner, error) {
	return srv.Repo.Get(Banner)
}

func (srv *BannerService) Edit(Banner model.Banner) (bool, error) {
	return srv.Repo.Edit(Banner)
}

func (srv *BannerService) Delete(id string) (bool, error) {
	return srv.Repo.Delete(id)
}

func (srv *BannerService) List(req *query.ListQuery) (Banners []*model.Banner, err error) {
	return srv.Repo.List(req)
}

func (srv *BannerService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *BannerService) Exist(Banner model.Banner) *model.Banner {
	return srv.Repo.Exist(Banner)
}

func (srv *BannerService) ExistByBannerID(id string) *model.Banner {
	return srv.Repo.ExistByBannerID(id)
}
