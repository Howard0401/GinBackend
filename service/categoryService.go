package service

import (
	"VueGin/model"
	"VueGin/repository"
	"VueGin/repository/query"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type CategorySrv interface {
	List(req *query.ListQuery) (Categories []*model.CategoryResult, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryId(id string) *model.Category
	Add(Category model.CategoryResult) (bool, error)
	Edit(Category model.Category) (bool, error)
	Delete(Category model.Category) (bool, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInterface
}

func (srv *CategoryService) ExistByCategoryId(id string) *model.Category {
	return srv.Repo.ExistByCategoryID(id)
}
func (srv *CategoryService) List(req *query.ListQuery) (Categories []*model.CategoryResult, err error) {
	return srv.Repo.List(req)
}

func (srv *CategoryService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *CategoryService) Get(id string) ([]*model.CategoryResult, error) {
	return srv.Repo.Get(id)
}

func (srv *CategoryService) Exist(Category model.Category) *model.Category {
	return srv.Repo.Exist(Category)
}

//因為有三級分類都在同張表中，因此需要先檢查分類名稱、
func (srv *CategoryService) Add(CategoryResult model.CategoryResult) (bool, error) {

	if CategoryResult.C1Name == "" || CategoryResult.C2Name == "" || CategoryResult.C3Name == "" {
		return false, fmt.Errorf("請輸入正確的類別名")
	}

	c1 := model.Category{
		CategoryId: CategoryResult.C1CategoryID,
		Name:       CategoryResult.C1Name,
		Desc:       CategoryResult.C1Desc,
		Order:      CategoryResult.C1Order,
		ParentId:   CategoryResult.C1ParentId,
	}
	r1 := srv.Exist(c1)

	c2 := model.Category{
		CategoryId: CategoryResult.C2CategoryID,
		Name:       CategoryResult.C2Name,
		Desc:       "",
		Order:      CategoryResult.C2Order,
		ParentId:   CategoryResult.C1CategoryID,
		IsDeleted:  false,
	}
	r2 := srv.Exist(c2)

	c3 := model.Category{
		CategoryId: CategoryResult.C3CategoryID,
		Name:       CategoryResult.C3Name,
		Desc:       "",
		Order:      CategoryResult.C3Order,
		ParentId:   CategoryResult.C2CategoryID,
		IsDeleted:  false,
	}
	r3 := srv.Exist(c3)
	//先檢查有沒有被添加過分類
	if r1.Name == c1.Name && r2.Name == c2.Name && r3.Name == c3.Name {
		return false, fmt.Errorf("分類已存在")
	}
	//先建立好分類名稱
	if c1.CategoryId == "" {
		c1ID := uuid.NewV4().String()
		c1.ParentId = uuid.NewV4().String()
		c1.CategoryId = c1ID
		c2.ParentId = c1ID
	}
	if c2.CategoryId == "" {
		c2ID := uuid.NewV4().String()
		c2.CategoryId = c2ID
		c3.ParentId = c2.CategoryId
	}
	if c3.CategoryId == "" {
		c3ID := uuid.NewV4().String()
		c3.CategoryId = c3ID
	}
	//若發現1.2層級分類相同時，僅添加第3層分類
	if r1.Name == c1.Name && r2.Name == c2.Name {
		c3.ParentId = r2.CategoryId
		_, err := srv.Repo.Add(c3)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	//若發現1層級分類相同時，僅添加第2.3層分類
	if r1.Name == c1.Name {
		c1.ParentId = r1.ParentId
		c1.CategoryId = r1.CategoryId
		c2.ParentId = r1.CategoryId
		_, err := srv.Repo.Add(c2)
		if err != nil {
			return false, err
		}
		_, err = srv.Repo.Add(c3)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	_, _ = srv.Repo.Add(c1)
	_, _ = srv.Repo.Add(c2)
	_, _ = srv.Repo.Add(c3)

	return true, nil
}

func (srv *CategoryService) Edit(Category model.Category) (bool, error) {
	return srv.Repo.Edit(Category)
}

func (srv *CategoryService) Delete(Category model.Category) (bool, error) {
	if Category.CategoryId == "" {
		return false, fmt.Errorf("參數錯誤")
	}
	return srv.Repo.Delete(Category)
}
