package handler

import (
	format "VueGin/Utils/logFormat"
	"VueGin/enum"
	"VueGin/model"
	"VueGin/repository/query"
	res "VueGin/resViewModel"
	"VueGin/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductSrv service.ProductSrv
}

func (h *ProductHandler) GetEntity(in model.Product) res.Product {
	return res.Product{
		Id:                   in.ProductId,
		Key:                  in.ProductId,
		ProductId:            in.ProductId,
		ProductName:          in.ProductName,
		ProductIntro:         in.ProductIntro,
		CategoryId:           in.CategoryId,
		CategoryName:         in.CategoryName,
		ProductCoverImg:      in.ProductCoverImg,
		ProductBanner:        in.ProductBanner,
		OriginalPrice:        in.OriginalPrice,
		SellingPrice:         in.SellingPrice,
		StockNum:             in.StockNum,
		Tag:                  in.Tag,
		SellStatus:           in.SellStatus,
		ProductDetailContent: in.ProductDetailContent,
		IsDeleted:            in.IsDeleted,
	}
}

func (h *ProductHandler) ProductInfo(c *gin.Context) {
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	pid := c.Param("id")
	if pid == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	input := model.Product{
		ProductId: pid,
	}
	result, err := h.ProductSrv.Get(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	data := h.GetEntity(*result)

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      data,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

func (h *ProductHandler) ProductList(c *gin.Context) {
	var q query.ListQuery
	entity := res.Entity{
		Code:      int(enum.ResOK),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	//因為要查詢的型別是 *query.ListQuery
	//context傳來的JSON綁定q
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	//確定分頁數量
	list, err := h.ProductSrv.List(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("List Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	total, err := h.ProductSrv.GetTotal(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("GetTotal Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	if q.PageSize == 0 {
		q.PageSize = 5
	}

	retNum := int(int(total) % q.PageSize)
	ret := int(int(total) / q.PageSize)

	pages := 0

	if retNum == 0 {
		pages = ret
	} else {
		pages = ret + 1
	}

	var outputList []*res.Product
	for _, item := range list {
		r := h.GetEntity(*item)
		outputList = append(outputList, &r)
	}

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: pages,
		Data:      outputList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	p := model.Product{}
	err := c.ShouldBindJSON(&p)
	// fmt.Println(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	result, err := h.ProductSrv.Add(p)
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	entity.Code = int(enum.ResOK)
	entity.Msg = "OK,已添加Data"
	entity.Data = result
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

func (h *ProductHandler) EditProduct(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	var p model.Product
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	success, err := h.ProductSrv.Edit(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = enum.ResOK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	success, err := h.ProductSrv.Delete(id)
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}

	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = fmt.Sprintf("刪除成功，id:%s", id)
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}
}
