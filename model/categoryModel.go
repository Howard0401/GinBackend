package model

type Category struct {
	CategoryId string `json:"categoryId" gorm:"column:category_id"`
	Name       string `json:"name" gorm:"column:name"`
	Desc       string `json:"desc" gorm:"column:desc"`
	Order      int    `json:"order" gorm:"column:order"`
	ParentId   string `json:"parentId" gorm:"column:parent_id"`
	IsDeleted  bool   `josn:"isDeleted" gorm:"column:is_deleted"`
	//ParentId 直接在一張表裡分級，不用關聯
}

//三級分類 大類(第一級) 每個分類(CategoryID)都有尋找下一層對應的子分類(ParentId)
type CategoryResult struct {
	C1CategoryID string `json:"c1CategoryId" gorm:"c1_category_id"`
	C1Name       string `json:"c1Name" gorm:"column:c1_name"`
	C1Desc       string `json:"c1Desc" gorm:"column:c1_desc"`
	C1Order      int    `json:"c1Order" gorm:"column:c1_order"`
	C1ParentId   string `json:"c1ParentId" gorm:"column:c1_parent_id"`

	C2CategoryID string `json:"c2CategoryId" gorm:"column:c2_category_id"`
	C2Name       string `json:"c2Name" gorm:"column:c2_name"`
	C2Order      int    `json:"c2Order" gorm:"column:c2_order"`
	C2ParentId   string `json:"c2ParentId" gorm:"column:c2_parent_id"`

	Key          string `json:"key"`
	Id           string `json:"id"`
	C3CategoryID string `json:"c3CategoryId" gorm:"c3_category_id"`
	C3Name       string `json:"c3Name" gorm:"column:c3_name"`
	C3Order      int    `json:"c3Order" gorm:"column:c3_order"`
	C3ParentId   string `json:"c3ParentId" gorm:"c3_parent_id"`
	IsDeleted    bool   `json:"isDeleted" gorm:"c3_is_deleted" `
}
