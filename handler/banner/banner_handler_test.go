package bannerhandler

import (
	RT "VueGin/Utils/for_router_test"
	"VueGin/global"
	"VueGin/middleware"
	"VueGin/model"
	"VueGin/repository"
	"VueGin/service"
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

var methods BannerHandler

func GetMethods() {
	methods = BannerHandler{
		BannerSrv: &service.BannerService{
			Repo: &repository.BannerRepository{
				DB: global.Global_DB,
			},
		}}
}

var r *gin.Engine

func LoadSettings() {
	RT.InitFor()
	GetMethods()
	r = gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(gin.TestMode)
}

func TestAddBanner(t *testing.T) {
	LoadSettings() //為什麼不用Init 因為不會預先載入 原因?=>待解
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
		log.Println("測試/api/banner/add 成功...")
		global.Global_Logger.Info("Test Success: /api/banner/add", zap.String("test_result", "success"))
	} else {
		t.Error("測試/api/banner/add 失敗...")
		global.Global_Logger.Debug("Test Failed: /api/banner/add", zap.String("test_result", string(body)))
	}
}

func TestAddBannerError(t *testing.T) {
	LoadSettings()
	//假設重復Url
	jsonStrDuplicateUrl := []byte(`
	{
  "url": "banner1",
  "redirectUrl": "redirectUrlTEST",
  "order": 1,
  "createUser": "addUser",
  "updateUser": "admin"
	}`)
	//假設JSON格式不符
	jsonStrWrongFormat := []byte(`
	{
	"bannerIdd":"strr",,,
  "url": "banner1",
  "redirectUrl": "redirectUrlTEST",
  "order": 1,
  "createUser": "addUser",
  "updateUser": "admin"
	}`)
	list := [][]byte{jsonStrWrongFormat, jsonStrDuplicateUrl}
	r.POST("/api/banner/add", methods.AddBanner)
	for i := range list {
		req, err := http.NewRequest("POST", "/api/banner/add", bytes.NewBuffer(list[i]))
		if err != nil {
			t.Error("測試/api/banner/add (Error JSON and Parameter) 失敗...")
		}
		newRecorder := httptest.NewRecorder()
		r.ServeHTTP(newRecorder, req)

		res := newRecorder.Result()
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Error("測試/api/banner/add (Error JSON and Parameter) 失敗...")
		}

		if newRecorder.Result().StatusCode == 200 {
			t.Error("測試/api/banner/add (Error JSON and Parameter) 失敗...")
			global.Global_Logger.Info("Test Success: /api/banner/add", zap.String("test_result", "success"))
		} else {
			log.Println("測試/api/banner/add (Error JSON and Parameter) 成功...")
			global.Global_Logger.Debug("Test Failed: /api/banner/add", zap.String("test_result", string(body)))
		}
	}
}

func TestBannerList(t *testing.T) {
	LoadSettings()
	r.GET("/api/banner/list", methods.BannerList)

	q := [][]string{
		{"1", "2"},
		{"0", "0"},
	}
	for i := 0; i < len(q); i++ {
		req, err := http.NewRequest("GET", "/api/banner/list", nil)
		if err != nil {
			t.Fatal(err)
		}
		query := req.URL.Query()
		q0 := q[i][0]
		q1 := q[i][1]

		query.Add("page", q0)
		query.Add("size", q1)

		req.URL.RawQuery = query.Encode()
		newRecorder := httptest.NewRecorder()
		r.ServeHTTP(newRecorder, req)
		//顯示查詢結果
		// fmt.Println(newRecorder.Body.String())
		if newRecorder.Result().StatusCode != 200 {
			t.Errorf("測試/api/banner/list?page=%v&size=%v 失敗!!!", q0, q1)
			global.Global_Logger.Debug("Test Failed /api/banner/list", zap.Int("test_result_code", newRecorder.Result().StatusCode))
		} else {
			log.Printf("測試/api/banner/list?page=%v&size=%v 成功!!!", q0, q1)
			global.Global_Logger.Info("Test Success: /api/banner/list", zap.String("test_result", "success"))
		}
		query.Del("page")
		query.Del("size")

	}
}

func TestBannerError(t *testing.T) {
	LoadSettings()
	r.GET("/api/banner/list", methods.BannerList)

	q := [][]string{
		{"ffftttt", "hhsss"},
		{"1", "sss"},
		{"sss", "111"},
		{"0", "fhgtrhtrh++23/"},
	}
	for i := range q {
		req, err := http.NewRequest("GET", "/api/banner/list", nil)
		if err != nil {
			t.Fatal(err)
		}
		query := req.URL.Query()
		query.Add("page", q[i][0])
		query.Add("size", q[i][1])

		req.URL.RawQuery = query.Encode()
		newRecorder := httptest.NewRecorder()
		r.ServeHTTP(newRecorder, req)
		//顯示查詢結果
		fmt.Println(newRecorder.Body.String())
		if newRecorder.Result().StatusCode != 200 {
			log.Printf("測試/api/banner/list?page=%v&size=%v 成功!!!", q[i][0], q[i][1])
			global.Global_Logger.Info("Test Success: /api/banner/list (Error Case)", zap.String("test_result", "success"))

		} else {
			t.Errorf("測試/api/banner/list?page=%v&size=%v 失敗!!!", q[i][0], q[i][1])
			global.Global_Logger.Debug("Test Failed /api/banner/list (Error Case)", zap.Int("test_result_code", newRecorder.Result().StatusCode))
		}
		query.Del("page")
		query.Del("size")
	}

}

func TestGetBannerInfo(t *testing.T) {
	LoadSettings()
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
		log.Println("測試 /api/banner/info?id= 成功...")
		global.Global_Logger.Info("Test Success: /api/banner/info?id= ", zap.String("test_result", "success"))
	} else {
		t.Error("測試 /api/banner/info?id=失敗!!!")
		global.Global_Logger.Debug("Test Failed: /api/banner/info?id=", zap.String("test_result", string(body)))
	}
}

func TestGetBannerInfoNilorErrorJSON(t *testing.T) {
	LoadSettings()
	r.GET("/api/banner/info", methods.BannerInfo)
	req, err := http.NewRequest("GET", "/api/banner/info", nil)
	if err != nil {
		t.Fatal(err)
	}
	query := req.URL.Query()
	//這邊很重要，因為從資料庫裡用url找到bannerId，再賦值給全域變數

	list := []string{
		"", "wrong test case",
	}
	// flag := true

	for i := 0; i < len(list); i++ {
		query.Add("id", list[i])
		req.URL.RawQuery = query.Encode()
		newRecorder := httptest.NewRecorder()
		r.ServeHTTP(newRecorder, req)
		res := newRecorder.Result()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		//絕對不能有空格，連一個都不能有！！！
		expected := fmt.Sprintf(`{"entity":{"code":200,"msg":"OK，查詢成功","total":0,"totalPage":0,"data":{"id":"%s","key":"%s","bannerId":"%s","url":"%s","redirectUrl":"redirectUrlTEST","order":1}}}`, idForTest, idForTest, idForTest, urlForTest)
		// fmt.Println(string(body))
		if !assert.IsEqual(string(body), expected) {
			log.Println("測試 /api/banner/info?id=nil or err => error 成功... (Error Case)")
			global.Global_Logger.Info("Test Success: /api/banner/info?id=nil or err => error ", zap.String("test_result", "success"))
		} else {
			// log.Println("測試/api/banner/info?id= 失敗!!!")
			t.Error("測試 /api/banner/info?id=nil or err => error 失敗... (Error Case)")
			global.Global_Logger.Debug("Test Failed: /api/banner/info?id=", zap.String("test_result", string(body)))
		}
	}

}

func TestPostBannerEdit(t *testing.T) {
	LoadSettings()
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
	fmt.Println(string(body))
	expected := fmt.Sprintf(`{"entity":{"code":200,"msg":"修改成功,id:%s","total":0,"totalPage":0,"data":null}}`, idForTest)

	// fmt.Println(string(body))
	if assert.IsEqual(string(body), expected) {
		log.Println("測試/api/banner/edit 成功...")
		global.Global_Logger.Info("Test Success: /api/banner/edit ", zap.String("test_result", "success"))
	} else {
		t.Error("測試/api/banner/edit 失敗...")
		global.Global_Logger.Debug("Test Failed:/api/banner/edit", zap.String("test_result", string(body)))
		// log.Println(newRecorder.Body)
	}
}

func TestPostBannerEditError(t *testing.T) {
	LoadSettings()
	r.POST("/api/banner/edit", methods.EditBanner)

	wrongJSONFormat := fmt.Sprintf(`{"banner_id":"%s",,,,"url":"%s","redirectUrl": "redirectUrlTESTEdited","order":1,"createUser":"addUser",  "updateUser":"admin"}`, idForTest, urlForTestEdit)
	emptyBannerId := fmt.Sprintf(`{"banner_id":"","url":"%s","redirectUrl": "redirectUrlTESTEdited","order":1,"createUser":"addUser",  "updateUser":"admin"}`, urlForTestEdit)
	wrongBannerId := fmt.Sprintf(`{"banner_id":"testestest111111","url":"%s","redirectUrl": "redirectUrlTESTEdited","order":1,"createUser":"addUser",  "updateUser":"admin"}`, urlForTestEdit)

	list := []string{
		wrongJSONFormat, emptyBannerId, wrongBannerId,
	}
	for i := range list {
		tmp := []byte(list[i])
		req, err := http.NewRequest("POST", "/api/banner/edit", bytes.NewBuffer(tmp))
		if err != nil {
			t.Fatal(err)
		}
		newRecorder := httptest.NewRecorder()
		r.ServeHTTP(newRecorder, req)

		// fmt.Println(string(body))
		if assert.IsEqual(newRecorder.Result().StatusCode, 200) {
			t.Error("測試/api/banner/edit 失敗... (Error Case)")
			global.Global_Logger.Debug("Test Failed:/api/banner/edit (Error Case)", zap.Int("test_result", newRecorder.Result().StatusCode))
		} else {
			log.Printf("測試/api/banner/edit 成功... (Error Case)")
			global.Global_Logger.Info("Test Success: /api/banner/edit (Error Case) ", zap.String("test_result", "success"))
		}
	}

}

func TestPostBannerDelete(t *testing.T) {
	LoadSettings()

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
		t.Error("測試/api/banner/delete?id=" + idForTest + "失敗...")
		global.Global_Logger.Debug("Test Failed: /api/banner/delete?id=", zap.Int("test_result_code", newRecorder.Result().StatusCode))
		// log.Println(newRecorder.Body)
	}
}

func TestPostBannerDeleteError(t *testing.T) {
	LoadSettings()

	r.POST("/api/banner/delete", methods.DeleteBanner)
	req, err := http.NewRequest("POST", "/api/banner/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	newRecorder := httptest.NewRecorder()
	//帶查詢參數要這樣改
	query := req.URL.Query()
	list := []string{
		"22222223333333",
		"test",
		"where 1=1 delete *",
	}
	for i := 0; i < len(list); i++ {
		query.Add("id", list[i])
		req.URL.RawQuery = query.Encode()

		r.ServeHTTP(newRecorder, req)

		if newRecorder.Result().StatusCode == 200 {
			t.Error("測試/api/banner/delete?id=" + list[i] + "失敗... (Error Case)")
			global.Global_Logger.Info("Test Success: /api/banner/delete?id= (Error Case)", zap.String("test_result", "success"))
			// log.Println(newRecorder.Body)
		} else {
			log.Println("測試/api/banner/delete?id=" + list[i] + "成功... (Error Case)")
			global.Global_Logger.Debug("Test Failed: /api/banner/delete?id= (Error Case)", zap.Int("test_result_code", newRecorder.Result().StatusCode))
			// log.Println(newRecorder.Body)
		}
	}
}
