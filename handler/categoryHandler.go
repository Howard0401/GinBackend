package handler

import (
	utils "VueGin/Utils"
	"VueGin/enum"
	"VueGin/model"
	"VueGin/repository/query"
	res "VueGin/resViewModel"
	"VueGin/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategorySrv service.CategorySrv
}

//給後臺，可以直接編輯第三級分類 (Service層使用原生SQL)
func (h *CategoryHandler) CategoryListForBackend(c *gin.Context) {
	var q query.ListQuery
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	//綁定query
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	if q.PageSize == 0 {
		q.PageSize = 5
	}
	//讀取全部分類
	list, err := h.CategorySrv.List(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("List Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	//計算該分類總數
	total, err := h.CategorySrv.GetTotal(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("GetTotal Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	//傳給前端要開的分頁數
	for _, item := range list {
		item.Key = item.C3CategoryID
		item.Id = item.C3CategoryID
	}
	pages := 0
	if int(total)%q.PageSize == 0 {
		pages = int(int(total) / q.PageSize)
	} else {
		pages = int(int(total)/q.PageSize) + 1
	}
	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: pages,
		Data:      list,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

//解包
func (h *CategoryHandler) GetEntity(result []*model.CategoryResult) map[string]*res.Category {
	//key值是ID value是細項，從第三層開始解包，慢慢往上層傳
	c3 := make(map[string]*res.Category3)
	for _, item := range result {
		thirdCategory := &res.Category3{
			Id:         item.C3CategoryID,
			Key:        item.C3CategoryID,
			CategoryID: item.C3CategoryID,
			Name:       item.C3Name,
			Order:      item.C3Order,
			ParentID:   item.C3ParentId,
			IsDeleted:  item.IsDeleted,
		}
		c3[item.C3CategoryID] = thirdCategory
	}
	c2 := make(map[string]*res.Category2)
	for _, item := range result {
		secondCategory := &res.Category2{
			CategoryID: item.C2CategoryID,
			Name:       item.C2Name,
			Order:      item.C2Order,
			ParentID:   item.C2ParentId,
			Children:   c3,
		}
		c2[item.C2CategoryID] = secondCategory
	}

	c1 := make(map[string]*res.Category)
	for _, item := range result {
		firstCategory := &res.Category{
			CategoryID: item.C1CategoryID,
			Name:       item.C1CategoryID,
			Order:      item.C1Order,
			Children:   c2,
		}
		c1[item.C1CategoryID] = firstCategory
	}
	return c1
}

//給前端
func (h *CategoryHandler) CategoryList(c *gin.Context) {
	var q query.ListQuery
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	if q.PageSize == 0 {
		q.PageSize = 5
	}
	list, err := h.CategorySrv.List(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("List Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	total, err := h.CategorySrv.GetTotal(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("GetTotal Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	//解包
	unpackMap := h.GetEntity(list)
	pages := 0
	if int(total)%q.PageSize == 0 {
		pages = int(int(total) / q.PageSize)
	} else {
		pages = int(int(total)/q.PageSize) + 1
	}

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: pages,
		Data:      unpackMap,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

func (h *CategoryHandler) CategoryInfo(c *gin.Context) {
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	cid := c.Param("id")
	if cid == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	result, err := h.CategorySrv.Get(cid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	unpack := h.GetEntity(result)

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      unpack,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

//增加子分類(model.CategoryResult{})
func (h *CategoryHandler) AddCategory(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	category := model.CategoryResult{}
	err := c.ShouldBindJSON(&category)
	fmt.Println(category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	success, err := h.CategorySrv.Add(category)
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	if !success {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	entity.Code = int(enum.ResOK)
	entity.Msg = fmt.Sprintf("添加分類成功:%v", category)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

//編輯分類
func (h *CategoryHandler) EditCategory(c *gin.Context) {
	category := model.Category{}
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}

	err := c.ShouldBindJSON(&category)
	fmt.Println(category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	fmt.Println(category)
	success, err := h.CategorySrv.Edit(category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = enum.ResOK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
	}
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	chkExistModel := h.CategorySrv.ExistByCategoryId(id)
	success, err := h.CategorySrv.Delete(*chkExistModel)
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = fmt.Sprintf("刪除成功,id:%s", id)
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
	}
}
