package repository

import (
	utils "VueGin/Utils"
	"VueGin/model"
	"VueGin/repository/query"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BannerRepository struct {
	DB *gorm.DB
}

type BannerRepoInterface interface {
	Exist(Banner model.Banner) *model.Banner
	Add(Banner model.Banner) (*model.Banner, error)
	Get(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(id string) (bool, error)

	List(req *query.ListQuery) (Banners []*model.Banner, err error)
	ExistByBannerID(id string) *model.Banner
	GetTotal(req *query.ListQuery) (total int64, err error) //總數
}

func (repo *BannerRepository) Exist(Banner model.Banner) *model.Banner {
	if Banner.Url != "" && Banner.RedirectUrl != "" {
		var b model.Banner
		repo.DB.Model(&Banner).Where("url=? and redirect_url", Banner.Url, Banner.RedirectUrl).First(&b)
		return &b
	}
	return nil
}

func (repo *BannerRepository) Add(Banner model.Banner) (*model.Banner, error) {
	// if Banner.BannerId == "" {
	// 	Banner.BannerId = uuid.NewV4().String()
	// }
	Banner.BannerId = uuid.NewV4().String()
	exist := repo.Exist(Banner)
	if exist != nil && exist.Url == Banner.Url && exist.RedirectUrl == Banner.RedirectUrl {
		return nil, fmt.Errorf("banner exist")
	}
	err := repo.DB.Create(Banner).Error
	if err != nil {
		return nil, fmt.Errorf("add banner failed:%v", err)
	}
	return &Banner, nil
}

func (repo *BannerRepository) Get(Banner model.Banner) (*model.Banner, error) {
	// 不用這樣寫 Find 已經包含 Where 語法
	// err := repo.DB.Find(&Banner).Where("banner_id=?", Banner.BannerId).Error
	err := repo.DB.Where("banner_id =  ?", Banner.BannerId).Find(&Banner).Error
	if err != nil {
		return nil, fmt.Errorf("get failed:%v", err)
	}
	// fmt.Println(&Banner)
	return &Banner, nil
}

func (repo *BannerRepository) Edit(Banner model.Banner) (bool, error) {
	if Banner.BannerId == "" {
		return false, fmt.Errorf("please input id")
	}
	b := &model.Banner{}
	err := repo.DB.Model(b).Where("banner_id = ? ", Banner.BannerId).Updates(map[string]interface{}{
		"url":          Banner.Url,
		"redirect_url": Banner.RedirectUrl,
		"order_by_idx": Banner.Order,
	}).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func (repo *BannerRepository) Delete(id string) (bool, error) {
	// err := repo.DB.Where("banner_id=?", id).Delete(&model.Banner{}).Error
	in := &model.Banner{BannerId: id}
	err := repo.DB.Where("banner_id=?", id).Delete(in).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *BannerRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var banners []model.Banner
	// db := repo.DB
	// if req.Where != "" {
	// 	db = db.Where(req.Where)
	// }
	if err := repo.DB.Find(&banners).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *BannerRepository) List(req *query.ListQuery) (banners []*model.Banner, err error) {

	db := repo.DB
	limit, offset := utils.Page(req.PageSize, req.Page)
	if err := db.Limit(limit).Offset(offset).Order("order_by_idx").Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

func (repo *BannerRepository) ExistByBannerID(id string) *model.Banner {
	var b model.Banner
	if err := repo.DB.Where("order_id=?", id).First(&b).Error; err != nil {
		return nil
	}
	return &b
}
