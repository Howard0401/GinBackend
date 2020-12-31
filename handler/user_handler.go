package handler

import (
	format "VueGin/Utils/logFormat"
	enum "VueGin/enum"
	model "VueGin/model"
	"VueGin/repository/query"
	res "VueGin/resViewModel"
	service "VueGin/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

//建立Entity 把數據隔離，讓以後需求變動時容易擴充
func (h *UserHandler) GetEntity(in model.User) res.User {
	return res.User{
		Id:        in.UserId,
		Key:       in.UserId,
		UserId:    in.UserId,
		NickName:  in.NickName,
		Mobile:    in.Mobile,
		Address:   in.Address,
		IsDeleted: in.IsDeleted,
		IsLocked:  in.IsLocked,
	}
}

func (h *UserHandler) UserInfo(c *gin.Context) {
	//create format of return value  先建立好輸出的實體格式
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0, //沒分頁
		TotalPage: 1,
		Data:      nil,
	}
	//從傳來的上下文包含的id，利用商業邏輯srv查詢User的資訊 get context with id, and use service to get user info
	uid := c.Param("id")
	if uid == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	query := model.User{
		UserId: uid,
	}
	result, err := h.UserSrv.Get(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	output := h.GetEntity(*result)
	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      output,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

func (h *UserHandler) UserList(c *gin.Context) {
	var q query.ListQuery
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 0,
		Data:      nil,
	}

	err := c.ShouldBindQuery(&q)
	// fmt.Println(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	//計算分頁數傳給前端
	list, err1 := h.UserSrv.List(&q)
	total, err2 := h.UserSrv.GetTotal(&q)
	fmt.Println(list)
	fmt.Println(total)
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
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
	var newList []*res.User
	for _, item := range list {
		r := h.GetEntity(*item)
		newList = append(newList, &r)
	}
	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: pages,
		Data:      newList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

func (h *UserHandler) AddUser(c *gin.Context) {
	//建立Entity
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	//與JSON格式做bind，傳入要加入的model
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	// fmt.Println(u)
	if err != nil {
		entity.Data = err
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	//使用商業邏輯Service(內含Repository)
	result, err := h.UserSrv.Add(u)

	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	//若加入ID為空
	if result.UserId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	//回傳
	entity.Code = int(enum.ResOK)
	entity.Msg = enum.ResOK.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

func (h *UserHandler) EditUser(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	success, err := h.UserSrv.Edit(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = enum.ResOK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	id := c.Param("id")
	success, err := h.UserSrv.Delete(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = enum.ResOK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}
}
