package banner

import (
	RT "VueGin/Utils/for_router_test"
	"VueGin/global"
	"VueGin/middleware"
	"VueGin/model"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/zap"
)

var (
	idForTest string //因為生成為uuid，所以需要其他欄位(Url 已經設為Unique)來找這個變數
)

const (
	urlForTest     = "inputBanner6"
	urlForTestEdit = "urlEdited"
)

func TestAddBanner(t *testing.T) {
	RT.InitFor()
	GetMethods()
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(gin.TestMode)

	jsonStr := []byte(`
	{
  "url": "inputBanner6",
  "redirectUrl": "redirectUrlTEST",
  "order": 1,
  "createUser": "addUser",
  "updateUser": "admin"
	}`)
	// fmt.Println(global.Global_DB)
	r.POST("/api/banner/add", methods.AddBanner)
	req, err := http.NewRequest("POST", "/api/banner/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
		t.Errorf("")
	}
	newRecorder := httptest.NewRecorder()
	r.ServeHTTP(newRecorder, req)

	res := newRecorder.Result()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("測試/api/banner/add 失敗...")
		t.Errorf("原因: %v", err)
	}

	if newRecorder.Result().StatusCode == 200 {
		global.Global_Logger.Info("Test Success: /api/banner/add", zap.String("test_result", "success"))
	} else {
		t.Error("測試/api/banner/add 失敗...")
		t.Errorf("得回傳值: %v", string(body))
		global.Global_Logger.Debug("Test Failed: /api/banner/add", zap.String("test_result", string(body)))
	}
}

func TestBannerList(t *testing.T) {
	RT.InitFor()
	GetMethods()
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(gin.TestMode)
	r.GET("/api/banner/list", methods.BannerList)
	req, err := http.NewRequest("GET", "/api/banner/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	newRecorder := httptest.NewRecorder()
	r.ServeHTTP(newRecorder, req)
	//顯示查詢結果
	// fmt.Println(newRecorder.Body.String())
	if newRecorder.Result().StatusCode != 200 {
		t.Error("測試/api/banner/list 失敗!!!")
		global.Global_Logger.Debug("Test Failed /api/banner/list", zap.Int("test_result_code", newRecorder.Result().StatusCode))
	} else {
		global.Global_Logger.Info("Test Success: /api/banner/list", zap.String("test_result", "success"))
	}
}

func TestGetBannerInfo(t *testing.T) {
	RT.InitFor()
	GetMethods()
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(gin.TestMode)

	r.GET("/api/banner/info", methods.BannerInfo)
	req, err := http.NewRequest("GET", "/api/banner/info", nil)
	if err != nil {
		t.Fatal(err)
	}
	//將進入的id加入URL中，不過這邊因為id為產生UUID存進資料庫，因此
	query := req.URL.Query()
	bannerModel, err := methods.BannerSrv.GetByUrl(model.Banner{Url: urlForTest})
	if err != nil {
		log.Printf("測試/api/banner/info?id= 失敗!!! 錯誤訊息: %v", err)
	}
	//這邊很重要，因為從資料庫裡用url找到bannerId，再賦值給全域變數
	idForTest = bannerModel.BannerId
	query.Add("id", idForTest)
	req.URL.RawQuery = query.Encode()

	newRecorder := httptest.NewRecorder()
	r.ServeHTTP(newRecorder, req)
	res := newRecorder.Result()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(body))
	// fmt.Println(idForTest)
	// fmt.Println(urlForTest)
	//絕對不能有空格，連一個都不能有！！！
	expected := fmt.Sprintf(`{"entity":{"code":200,"msg":"OK，查詢成功","total":0,"totalPage":0,"data":{"id":"%s","key":"%s","bannerId":"%s","url":"%s","redirectUrl":"redirectUrlTEST","order":1}}}`, idForTest, idForTest, idForTest, urlForTest)
	// fmt.Println(string(body))
	if assert.IsEqual(string(body), expected) {
		global.Global_Logger.Info("Test Success: /api/banner/info?id= ", zap.String("test_result", "success"))
	} else {
		log.Println("測試/api/banner/info?id= 失敗!!!")
		global.Global_Logger.Debug("Test Failed: /api/banner/info?id=", zap.String("test_result", string(body)))
	}
}

func TestPostBannerEdit(t *testing.T) {
	RT.InitFor()
	GetMethods()
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(gin.TestMode)
	tmp := fmt.Sprintf(`{"banner_id":"%s","url":"%s","redirectUrl": "redirectUrlTESTEdited","order":1,"createUser":"addUser",  "updateUser":"admin"}`, idForTest, urlForTestEdit)
	jsonStr := []byte(tmp)
	r.POST("/api/banner/edit", methods.EditBanner)
	req, err := http.NewRequest("POST", "/api/banner/edit", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	newRecorder := httptest.NewRecorder()
	r.ServeHTTP(newRecorder, req)
	res := newRecorder.Result()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(body))
	expected := fmt.Sprintf(`{"entity":{"code":200,"msg":"修改成功,id:%s","total":0,"totalPage":0,"data":null}}`, idForTest)

	// fmt.Println(string(body))
	if assert.IsEqual(string(body), expected) {
		log.Println("測試/api/banner/edit 成功...")
		global.Global_Logger.Info("Test Success: /api/banner/edit ", zap.String("test_result", "success"))
	} else {
		log.Println("測試/api/banner/edit 失敗...")
		global.Global_Logger.Debug("Test Failed:/api/banner/edit", zap.String("test_result", string(body)))
		// log.Println(newRecorder.Body)
	}
}

func TestPostBannerDelete(t *testing.T) {
	RT.InitFor()
	GetMethods()
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(gin.TestMode)

	r.POST("/api/banner/delete", methods.DeleteBanner)
	req, err := http.NewRequest("POST", "/api/banner/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	newRecorder := httptest.NewRecorder()
	//帶查詢參數要這樣改
	query := req.URL.Query()
	query.Add("id", idForTest)
	req.URL.RawQuery = query.Encode()

	r.ServeHTTP(newRecorder, req)

	if newRecorder.Result().StatusCode == 200 {
		log.Println("測試/api/banner/delete?id=" + idForTest + "成功...")
		global.Global_Logger.Info("Test Success: /api/banner/delete?id= ", zap.String("test_result", "success"))
		// log.Println(newRecorder.Body)
	} else {
		log.Println("測試/api/banner/delete?id=" + idForTest + "失敗...")
		global.Global_Logger.Debug("Test Failed: /api/banner/delete?id=", zap.Int("test_result_code", newRecorder.Result().StatusCode))
		// log.Println(newRecorder.Body)
	}
}
