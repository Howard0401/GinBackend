package handler

import (
	format "VueGin/Utils/logFormat"
	"VueGin/enum"
	"VueGin/model"
	res "VueGin/model/res_view_model"
	"VueGin/repository/query"
	"VueGin/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//引入介面
type BannerHandler struct {
	BannerSrv service.BannerSrv
}

func (h *BannerHandler) GetEntity(in model.Banner) res.Banner {
	return res.Banner{
		Id:          in.BannerId,
		Key:         in.BannerId,
		BannerId:    in.BannerId,
		Url:         in.Url,
		RedirectUrl: in.RedirectUrl,
		OrderBy:     in.Order,
	}
}

// @Summary Banner Add
// @Tags Banner
// @Produce  json
// @Param id query string true "Info Query id"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "請求錯誤"
// @Failure 500 {string} string "內部錯誤"
// @Router /api/banner/info [get]
func (h *BannerHandler) BannerInfo(c *gin.Context) {
	// bid := c.Param("id") //這個是/info/:id的路由 (api/banner/info/inputBanner1)
	bid := c.Query("id") // (api/banner/info?id=inputBanner1)
	entity := res.Entity{
		Code:      int(enum.ResFail),
		Msg:       enum.ResFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}

	// fmt.Println(bid)
	if bid == "" {
		entity.Msg = "請傳入ID!!!"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	b := model.Banner{BannerId: bid} //若不指定id，會回傳資料庫中的第一筆
	result, err := h.BannerSrv.Get(b)
	// fmt.Println(result)
	if err != nil {
		entity.Msg = fmt.Sprintf("查詢時發生錯誤:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	//解包，調整為輸出給前端的格式(Entity)
	r := h.GetEntity(*result)
	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK，查詢成功",
		Total:     0,
		TotalPage: 0,
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

// @Summary Banner List
// @Tags Banner
// @Produce  json
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "請求錯誤"
// @Failure 500 {string} string "內部錯誤"
// @Router /api/banner/list [get]
func (h *BannerHandler) BannerList(c *gin.Context) {

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
		entity.Msg = fmt.Sprintf("提醒後端傳入query錯誤: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	list, err := h.BannerSrv.List(&q)

	if err != nil {
		entity.Msg = fmt.Sprintf("List Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	total, err := h.BannerSrv.GetTotal(&q)

	if err != nil {
		entity.Msg = fmt.Sprintf("GetTotal Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
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

	var result []*res.Banner
	//查詢結果為list，每迴圈都讀取list中第n項結構體，並調整為需要的輸出格式(Entity)
	for _, n := range list {
		tmp := h.GetEntity(*n)
		//取值，append進
		result = append(result, &tmp)
	}

	entity = res.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: pages,
		Data:      result,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)

}

// @Summary Banner Add
// @Tags Banner
// @Produce  json
// @Param b body model.Banner true "Add Banner model"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "請求錯誤"
// @Failure 500 {string} string "內部錯誤"
// @Router /api/banner/add [post]
func (h *BannerHandler) AddBanner(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	b := model.Banner{}
	err := c.ShouldBindJSON(&b)
	if err != nil {
		entity.Msg = fmt.Sprintf("傳入JSON錯誤: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	result, err := h.BannerSrv.Add(b)

	if err != nil {
		entity.Msg = fmt.Sprintf("Add Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}

	entity.Code = int(enum.ResOK)
	entity.Msg = fmt.Sprintf("添加成功:%v", result)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	format.EntityLog(entity)
}

// @Summary Banner Edit
// @Tags Banner
// @Produce  json
// @Param b body model.Banner true "Edit Banner model"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "請求錯誤"
// @Failure 500 {string} string "內部錯誤"
// @Router /api/banner/edit [post]
func (h *BannerHandler) EditBanner(c *gin.Context) {
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}

	b := model.Banner{}
	err := c.ShouldBindJSON(&b)
	// fmt.Println(b)
	if err != nil {
		entity.Msg = fmt.Sprintf("傳入JSON錯誤: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	success, err := h.BannerSrv.Edit(b)
	if err != nil {
		entity.Msg = fmt.Sprintf("Edit Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		format.EntityLog(entity)
		return
	}
	if success {
		entity.Code = int(enum.ResOK)
		entity.Msg = fmt.Sprintf("修改成功,id:%v", b.BannerId)
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		format.EntityLog(entity)
	}
}

// @Summary Banner Add
// @Tags Banner
// @Produce  json
// @Param id query string true "Delete by id"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "請求錯誤"
// @Failure 500 {string} string "內部錯誤"
// @Router /api/banner/delete [post]
func (h *BannerHandler) DeleteBanner(c *gin.Context) {
	// id := c.Param("id")
	id := c.Query("id")
	entity := res.Entity{
		Code:  int(enum.ResFail),
		Msg:   enum.ResFail.String(),
		Total: 0,
		Data:  nil,
	}
	// fmt.Println(id)
	success, err := h.BannerSrv.Delete(id)
	if err != nil {
		entity.Msg = fmt.Sprintf("Delete Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
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
