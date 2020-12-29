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

type OrderHandler struct {
	OrderSrv service.OrderSrv
}

func (h *OrderHandler) GetEntity(in model.Order) res.Order {
	return res.Order{
		Key:         in.OrderId,
		Id:          in.OrderId,
		OrderId:     in.OrderId,
		NickName:    in.NickName,
		Mobile:      in.Mobile,
		TotalPrice:  in.TotalPrice,
		PayStatus:   in.PayStatus,
		PayType:     in.PayType,
		PayTime:     in.PayTime,
		OrderStatus: in.OrderStatus,
		ExtraInfo:   in.ExtraInfo,
		UserAddress: in.UserAddress,
		IsDeleted:   in.IsDeleted,
	}
}

func (h *OrderHandler) OrderInfo(c *gin.Context) {
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	//Read id from context
	oid := c.Param("id")
	if oid == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	o := model.Order{
		OrderId: oid,
	}
	// fmt.Println(o)
	//從商業邏輯層套用Get方法，以ID查找符合之結果
	result, err := h.OrderSrv.Get(o)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	unpack := h.GetEntity(*result)

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      unpack,
	}
	c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

func (h *OrderHandler) OrderList(c *gin.Context) {
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
	list, err := h.OrderSrv.List(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("List Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	total, err := h.OrderSrv.GetTotal(&q)
	if err != nil {
		entity.Msg = fmt.Sprintf("GetTotal Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	if q.PageSize == 0 {
		q.PageSize = 5
	}

	retNum := int(int(total) / q.PageSize)
	ret := int(int(total) % q.PageSize)
	pages := 0
	if ret == 0 {
		pages = retNum
	} else {
		pages = retNum + 1
	}

	var output []*res.Order

	for _, n := range list {
		tmp := h.GetEntity(*n)
		output = append(output, &tmp)
	}

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: pages,
		Data:      output,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

func (h *OrderHandler) AddOrder(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	//加入新的order
	o := model.Order{}
	//綁定JSON，修改結構體中的o屬性
	err := c.ShouldBindJSON(&o)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	result, err := h.OrderSrv.Add(o)

	if err != nil {
		entity.Msg = err.Error()
		utils.EntityLog(entity)
		return
	}

	if result.OrderId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	entity.Code = int(enum.ResOK)
	entity.Msg = enum.ResOK.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	utils.EntityLog(entity)
}

func (h *OrderHandler) EditOrder(c *gin.Context) {
	o := model.Order{}
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&o)

	if err != nil || o.OrderId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	success, err := h.OrderSrv.Edit(o)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = fmt.Sprintf("編輯成功 %v", o)
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
	}
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	id := c.Param("id")
	o := h.OrderSrv.ExistByOrderId(id)
	//沒有返回o時
	if o == nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}
	success, err := h.OrderSrv.Delete(*o)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
		return
	}

	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = fmt.Sprintf("刪除成功, id:%s", id)
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		utils.EntityLog(entity)
	}
}
