package RT

import (
	"VueGin/config"
	global "VueGin/global"
	initSettings "VueGin/initSettings"
)

func InitFor() {
	//1.匯入Cnnfig設定 Import Viper to set up global config
	CongfigName := ""
	config, err := config.InitViper(CongfigName)
	if err != nil {
		panic(err)
	}
	// CongfigName = config.Name
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
}
