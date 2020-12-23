package repository

import (
	"VueGin/model"
	"VueGin/repository/query"
	"fmt"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req *query.ListQuery) (Categories []*model.CategoryResult, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(Category model.Category) (bool, error)
}

func (repo *CategoryRepository) List(req *query.ListQuery) (categories []*model.CategoryResult, err error) {
	var list []*model.CategoryResult
	err = repo.DB.Raw("SELECT c1.category_id as c1_category_id, c1.name as c1_name, c1.order as c1_order, c1.parent_id as c1_parent_id, c2.category_id as c2_category_id, c2.name as c2_name, c2.order as c2_order, c2.parent_id as c2_parent_id, c3.category_id as c3_category_id, c3.name as c3_name, c3.order as c3_order, c3.parent_id as c3_parent_id, c3.is_deleted as c3_is_deleted FROM category c1 join category c2 on c1.category_id = c2.parent_id join category c3 on c2.category_id = c3.parent_id").Find(&list).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(list); i++ {
		list[i].Key = list[i].C3CategoryID
	}
	return list, nil
}

func (repo *CategoryRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	err = repo.DB.Raw("SELECT count(c3.category_id) FROM category c1 join category c2 on c1.category_id = c2.parent_id join category c3 on c2.category_id=c3.parent_id").Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, err
}

func (repo *CategoryRepository) Get(id string) ([]*model.CategoryResult, error) {
	var list []*model.CategoryResult
	err := repo.DB.Raw("SELECT c1.category_id as c1_category_id,c1.name as c1_name,c1.desc as c1_desc,c1.order as c1_order,c2.category_id as c2_category_id,c2.name as c2_name,c2.order as c2_order,c3.category_id as c3_category_id,c3.name as c3_name,c3.order as c3_order FROM category c1 join category c2 on c1.category_id = c2.parent_id join category c3 on c2.category_id=c3.parent_id where c3.category_id = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *CategoryRepository) Exist(Category model.Category) *model.Category {
	var c model.Category
	if Category.Name != "" {
		repo.DB.Model(&c).Where("name=?", Category.Name).First(&c)
	}
	return &c
}
func (repo *CategoryRepository) ExistByCategoryID(id string) *model.Category {
	var c model.Category
	repo.DB.Where("category_id=?", id).First(&c)
	return &c
}

func (repo *CategoryRepository) Add(Category model.Category) (*model.Category, error) {
	err := repo.DB.Create(Category).Error
	if err != nil {
		return nil, fmt.Errorf("建立商品分類失敗")
	}
	return &Category, nil
}

func (repo *CategoryRepository) Edit(Category model.Category) (bool, error) {
	if Category.CategoryId == "" {
		return false, fmt.Errorf("請傳入ID")
	}
	c := &model.Category{}
	err := repo.DB.Model(c).Where("category_id=?", Category.CategoryId).Update(
		"name", Category.Name).Update("order", Category.Order).Error

	// err := repo.DB.Model(c).Where("category_id=?", Category.CategoryId).Updates(map[string]interface{}{
	// 	"name":  Category.Name,
	// 	"order": Category.Order,
	// }).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CategoryRepository) Delete(Category model.Category) (bool, error) {
	//軟刪除
	err := repo.DB.Model(&Category).Where("category_id=?", Category.CategoryId).Update("is_deleted", !Category.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
