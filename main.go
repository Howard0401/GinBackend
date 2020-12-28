package main

import (
	"VueGin/config"
	"VueGin/global"
	"VueGin/initSettings"
	server "VueGin/server"
)

var CongfigName string

func init() {
	//1.匯入Cnnfig設定 Import Viper to set up config
	CongfigName = ""
	config, err := config.InitViper(CongfigName)
	if err != nil {
		panic(err)
	}
	CongfigName = config.Name
	global.Global_Viper = config.Vp
	//2. 匯入GORM  Add Gorm to global variable
	global.Global_DB, err = initSettings.Gorm()
	if err != nil {
		panic(err)
	}

}
func main() {

	//1.匯入Cnnfig設定 Import Viper to set up config
	//2. 匯入GORM  Add Gorm to global variable
	// 以上使用 init()初始化(為什麼要這樣寫呢？=>為了unit test單獨run時，可以讀取到起始的DB)
	//https: //stackoverflow.com/questions/58913826/unit-test-code-runtime-error-invalid-memory-address-or-nil-pointer-dereferen
	// 所以，以後要先檢查DB有沒有初始化被讀取到(重點)
	db, err := global.Global_DB.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//3.全域已讀取Config與DB，執行Gin伺服器  Run Gin Server
	server.RunServer()
}
