package repository

import (
	utils "VueGin/Utils"
	"VueGin/model"
	query "VueGin/repository/query"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	Exist(Product model.Product) *model.Product
	Add(Product model.Product) (*model.Product, error)
	Get(Product model.Product) (*model.Product, error)
	Edit(Prodcut model.Product) (bool, error)
	Delete(p model.Product) (bool, error)

	List(req *query.ListQuery) (Products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	ExistByProductID(id string) (*model.Product, error)
}

//待改
func (repo *ProductRepository) ExistByProductID(id string) (*model.Product, error) {
	var p model.Product
	err := repo.DB.Model(&p).Where("product_id = ?", id).Find(&p).Error
	if err != nil {
		return &p, err
	}
	return &p, nil
}

func (repo *ProductRepository) List(req *query.ListQuery) (products []*model.Product, err error) {
	// fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页

	if err := db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

/////////////
func (repo *ProductRepository) Exist(Product model.Product) *model.Product {
	if Product.ProductName != "" {
		var tmp model.Product
		repo.DB.Where("product_name = ? ", Product.ProductName).First(&tmp)
		return &tmp
	}
	return nil
}

func (repo *ProductRepository) FindCategoryIdByName(name string) (string, error) {
	var n model.Category
	err := repo.DB.Model(&model.Category{}).Where("category.name= ?", name).First(&n).Error
	if err != nil {
		return "", err
	}
	return n.CategoryId, err
}

//Create
func (repo *ProductRepository) Add(Product model.Product) (*model.Product, error) {
	if exist := repo.Exist(Product); exist != nil && exist.ProductName != "" {
		return &Product, fmt.Errorf("商品已存在")
	}

	category_id, err := repo.FindCategoryIdByName(Product.CategoryName)

	if err != nil {
		return &Product, fmt.Errorf("新增產品時，判斷分類不存在: %v", err)
	}

	Product = model.Product{
		ProductId:            uuid.NewV4().String(),
		ProductName:          Product.ProductName,
		ProductIntro:         Product.ProductIntro,
		ProductCoverImg:      Product.ProductCoverImg,
		ProductBanner:        Product.ProductBanner,
		CategoryName:         Product.CategoryName,
		CategoryId:           category_id,
		OriginalPrice:        Product.OriginalPrice,
		SellingPrice:         Product.SellingPrice,
		StockNum:             Product.StockNum,
		Tag:                  Product.Tag,
		SellStatus:           Product.SellStatus,
		CreateUser:           "admin", //這邊測試時先寫死，之後如有分後台管理員ID，可以再調整傳入api的參數
		UpdateUser:           "admin",
		ProductDetailContent: Product.ProductDetailContent,
		IsDeleted:            false,
	}
	err = repo.DB.Create(Product).Error
	if err != nil {
		return nil, fmt.Errorf("商品添加失敗，原因：%v", err)
	}
	return &Product, nil
}

//Read
func (repo *ProductRepository) Get(Product model.Product) (*model.Product, error) {
	if err := repo.DB.Where(&Product).Find(&Product).Error; err != nil {
		return nil, err
	}
	return &Product, nil
}

//Update
func (repo *ProductRepository) Edit(Product model.Product) (bool, error) {
	if Product.ProductId == "" {
		return false, fmt.Errorf("請傳入更新ID")
	}
	// category_id, err := repo.FindCategoryIdByName(Product.CategoryName)
	// if err != nil {
	// 	return false, fmt.Errorf("請傳入商品名稱")
	// }
	p := &model.Product{}
	err := repo.DB.Model(p).Where("product_id=?", Product.ProductId).Updates(map[string]interface{}{
		"product_id":    Product.ProductId,
		"product_name":  Product.ProductName,
		"product_intro": Product.ProductIntro,
		"category_name": Product.CategoryName,
		// "category_id":            category_id,
		"category_id":            Product.CategoryId,
		"product_cover_img":      Product.ProductCoverImg,
		"product_banner":         Product.ProductBanner,
		"original_price":         Product.OriginalPrice,
		"selling_price":          Product.SellingPrice,
		"stock_num":              Product.StockNum,
		"tag":                    Product.Tag,
		"sell_status":            Product.SellStatus,
		"product_detail_content": Product.ProductDetailContent}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//Delete
func (repo *ProductRepository) Delete(Product model.Product) (bool, error) {
	err := repo.DB.Model(&Product).Where("product_id=?", Product.ProductId).Update("is_deleted", Product.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *ProductRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var products []*model.Product
	// db := repo.DB
	// if req.Where != "" {
	// 	db = db.Where(req.Where)
	// }
	if err := repo.DB.Find(&products).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}
