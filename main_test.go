package main

// import (
// 	"VueGin/global"
// 	"VueGin/handler"
// 	"VueGin/middleware"
// 	"VueGin/repository"
// 	"VueGin/service"
// 	"bytes"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

// // var (
// // 	BannerHandler   handler.BannerHandler
// // 	CategoryHandler handler.CategoryHandler
// // 	OrderHandler    handler.OrderHandler
// // 	ProductHandler  handler.ProductHandler
// // 	UserHandler     handler.UserHandler
// // 	idForTest       string //因為生成為uuid，所以需要其他欄位來找這個變數
// // 	// urlForTest      string
// // )

// const (
// 	urlForTest     = "inputBanner6"
// 	urlForTestEdit = "urlEdited"
// )

// func init() {
// 	initHandler()
// }

// func initHandler() {
// 	BannerHandler = handler.BannerHandler{
// 		BannerSrv: &service.BannerService{
// 			Repo: &repository.BannerRepository{
// 				DB: global.Global_DB,
// 			},
// 		}}

// 	CategoryHandler = handler.CategoryHandler{
// 		CategorySrv: &service.CategoryService{
// 			Repo: &repository.CategoryRepository{
// 				DB: global.Global_DB,
// 			},
// 		},
// 	}

// 	OrderHandler = handler.OrderHandler{
// 		OrderSrv: &service.OrderService{
// 			Repo: &repository.OrderRepository{
// 				DB: global.Global_DB,
// 			},
// 		}}

// 	ProductHandler = handler.ProductHandler{
// 		ProductSrv: &service.ProductService{
// 			Repo: &repository.ProductRepository{
// 				DB: global.Global_DB,
// 			},
// 		}}

// 	UserHandler = handler.UserHandler{
// 		UserSrv: &service.UserService{
// 			Repo: &repository.UserRepository{
// 				DB: global.Global_DB,
// 			},
// 		}}
// }

// func TestPostBannerAdd(t *testing.T) {
// 	r := gin.Default()
// 	r.Use(middleware.Cors())
// 	gin.SetMode(gin.TestMode)
// 	jsonStr := []byte(`
// 	{
//   "url": "inputBanner6",
//   "redirectUrl": "redirectUrlTEST",
//   "order": 1,
//   "createUser": "addUser",
//   "updateUser": "admin"
// 	}`)
// 	r.POST("/api/banner/add", BannerHandler.AddBanner)
// 	req, err := http.NewRequest("POST", "/api/banner/add", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	newRecorder := httptest.NewRecorder()
// 	r.ServeHTTP(newRecorder, req)

// 	if newRecorder.Result().StatusCode == 200 {
// 		log.Println("測試/api/banner/add 成功...")
// 		// log.Println(newRecorder.Body)
// 	} else {
// 		log.Println("測試/api/banner/add 失敗...")
// 		// log.Println(newRecorder.Body)
// 	}
// }

// // func TestGetBannerList(t *testing.T) {
// // 	r := gin.Default()
// // 	r.Use(middleware.Cors())
// // 	gin.SetMode(gin.TestMode)
// // 	r.GET("/api/banner/list", BannerHandler.BannerList)
// // 	req, err := http.NewRequest("GET", "/api/banner/list", nil)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	newRecorder := httptest.NewRecorder()
// // 	r.ServeHTTP(newRecorder, req)
// // 	//顯示查詢結果
// // 	// fmt.Println(newRecorder.Body.String())
// // 	if newRecorder.Result().StatusCode != 200 {
// // 		log.Println("測試/api/banner/list 失敗!!!")
// // 	} else {
// // 		log.Println("測試/api/banner/list成功...")
// // 	}
// // }

// // func TestGetBannerInfo(t *testing.T) {
// // 	r := gin.Default()
// // 	r.Use(middleware.Cors())
// // 	gin.SetMode(gin.TestMode)

// // 	r.GET("/api/banner/info", BannerHandler.BannerInfo)
// // 	req, err := http.NewRequest("GET", "/api/banner/info", nil)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	//將進入的id加入URL中，不過這邊因為id為產生UUID存進資料庫，因此
// // 	query := req.URL.Query()
// // 	bannerModel, err := BannerHandler.BannerSrv.GetByUrl(model.Banner{Url: urlForTest})
// // 	if err != nil {
// // 		log.Printf("測試/api/banner/info?id= 失敗!!! 錯誤訊息: %v", err)
// // 		return
// // 	}
// // 	//這邊很重要，因為從資料庫裡用url找到bannerId，再賦值給全域變數
// // 	idForTest = bannerModel.BannerId
// // 	query.Add("id", idForTest)
// // 	req.URL.RawQuery = query.Encode()

// // 	newRecorder := httptest.NewRecorder()
// // 	r.ServeHTTP(newRecorder, req)
// // 	res := newRecorder.Result()

// // 	body, err := ioutil.ReadAll(res.Body)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	// fmt.Println(string(body))
// // 	// fmt.Println(idForTest)
// // 	// fmt.Println(urlForTest)
// // 	//絕對不能有空格，連一個都不能有！！！
// // 	expected := fmt.Sprintf(`{"entity":{"code":200,"msg":"OK，查詢成功","total":0,"totalPage":0,"data":{"id":"%s","key":"%s","bannerId":"%s","url":"%s","redirectUrl":"redirectUrlTEST","order":1}}}`, idForTest, idForTest, idForTest, urlForTest)
// // 	// fmt.Println(string(body))
// // 	if string(body) != expected {
// // 		log.Println("測試/api/banner/info?id= 失敗!!!")
// // 	} else {
// // 		log.Println("測試/api/banner/info?id= 成功...")
// // 	}
// // }

// // func TestPostBannerEdit(t *testing.T) {
// // 	r := gin.Default()
// // 	r.Use(middleware.Cors())
// // 	gin.SetMode(gin.TestMode)
// // 	tmp := fmt.Sprintf(`{"bannerId":"%s","url":"%s","redirectUrl": "redirectUrlTESTEdited","order":1,"createUser":"addUser",  "updateUser":"admin"}`, idForTest, urlForTestEdit)
// // 	jsonStr := []byte(tmp)
// // 	r.POST("/api/banner/edit", BannerHandler.EditBanner)
// // 	req, err := http.NewRequest("POST", "/api/banner/edit", bytes.NewBuffer(jsonStr))
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	newRecorder := httptest.NewRecorder()
// // 	r.ServeHTTP(newRecorder, req)
// // 	res := newRecorder.Result()

// // 	body, err := ioutil.ReadAll(res.Body)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	expected := fmt.Sprintf(`{"entity":{"code":200,"msg":"OK，查詢成功","total":0,"totalPage":0,"data":{"id":"%s","key":"%s","bannerId":"%s","url":"%s","redirectUrl":"redirectUrlTEST","order":1}}}`, idForTest, idForTest, idForTest, urlForTestEdit)

// // 	// fmt.Println(string(body))
// // 	if string(body) != expected {
// // 		// if newRecorder.Result().StatusCode == 200 {
// // 		log.Println("測試/api/banner/edit 成功...")
// // 		// log.Println(newRecorder.Body)
// // 	} else {
// // 		log.Println("測試/api/banner/edit 失敗...")
// // 		// log.Println(newRecorder.Body)
// // 	}
// // }

// // func TestPostBannerDelete(t *testing.T) {
// // 	r := gin.Default()
// // 	r.Use(middleware.Cors())
// // 	gin.SetMode(gin.TestMode)

// // 	r.POST("/api/banner/delete", BannerHandler.DeleteBanner)
// // 	req, err := http.NewRequest("POST", "/api/banner/delete", nil)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	newRecorder := httptest.NewRecorder()
// // 	//帶查詢參數要這樣改
// // 	query := req.URL.Query()
// // 	query.Add("id", idForTest)
// // 	req.URL.RawQuery = query.Encode()

// // 	r.ServeHTTP(newRecorder, req)

// // 	if newRecorder.Result().StatusCode == 200 {
// // 		log.Println("測試/api/banner/delete?id=" + idForTest + "成功...")
// // 		// log.Println(newRecorder.Body)
// // 	} else {
// // 		log.Println("測試/api/banner/delete?id=" + idForTest + "失敗...")
// // 		// log.Println(newRecorder.Body)
// // 	}
// // }
