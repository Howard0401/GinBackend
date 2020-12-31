package main

import (
	"VueGin/Utils/tracer"
	"VueGin/config"
	"VueGin/global"
	"VueGin/initSettings"
	server "VueGin/server"
)

var CongfigName string = ""

func init() {
	//1.匯入Cnnfig設定 Import Viper to set up global config
	// CongfigName = ""
	config, err := config.InitViper(CongfigName)
	if err != nil {
		panic(err)
	}
	CongfigName = config.Name
	global.Global_Viper = config.Vp //雖已使用mapstructure，將.yaml檔匯入Config中，此處仍先暫時保留Viper

	//2. 匯入GORM  Add Gorm to global variable
	global.Global_DB, err = initSettings.Gorm()
	if err != nil {
		panic(err)
	}

	//3.初始化zap logger參數
	global.Global_Logger, err = initSettings.InitLogger()
	if err != nil {
		panic(err)
	}

	//4.初始化Jaeger參數 //未來或許可以加入Gorm追蹤，但目前還不成熟(Github還未有完整版V2)
	name := global.Global_Config.System.Name
	t, _, err := tracer.NewJaegerTracer(name, "127.0.0.1:6831")
	if err != nil {
		panic(err)
	}
	global.Global_Tracer = &t

}

// @title Gin-Backend
// @version 1.0
// @description golang作品
// @termsOfService github
func main() {
	// 以上使用 init()初始化(為什麼要這樣寫呢？=>為了unit test單獨run時，可以讀取到起始的DB)
	//https: //stackoverflow.com/questions/58913826/unit-test-code-runtime-error-invalid-memory-address-or-nil-pointer-dereferen
	// 所以，以後要先檢查DB有沒有初始化被讀取到(重點)

	db, err := global.Global_DB.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close() //主程式結束後關閉DB

	//5.全域已讀取Config與DB，執行Gin伺服器  Run Gin Server
	server.RunServer()
}
