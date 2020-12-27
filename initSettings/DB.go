package initSettings

import (
	"VueGin/global"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBconf struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Charset  string //編碼
}

func Gorm() (*gorm.DB, error) {
	// fmt.Println(global.Global_Viper)
	e := global.Global_Viper.GetString("system.first_db")
	fmt.Println(e)
	switch global.Global_Viper.GetString("system.first_db") {
	case "mysql":
		return GormMySQL()
	default:
		return nil, fmt.Errorf("未開啟資料庫...")
	}
}

func GormMySQL() (*gorm.DB, error) {
	fmt.Println("init MySQLDB starting...")
	dbStr := &DBconf{
		Host:     global.Global_Viper.GetString("database.host"),
		User:     global.Global_Viper.GetString("database.username"),
		Password: global.Global_Viper.GetString("database.password"),
		DBName:   global.Global_Viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		dbStr.User, dbStr.Password, dbStr.Host, dbStr.DBName, true, "Local",
	)
	fmt.Printf("%v\n", config)
	// gorm v1
	//DB, err = gorm.Open("mysql", config)
	//注意下方是gorm v2的版本
	//https://github.com/go-gorm/mysql
	//特別注意，這邊DB要注意起始值，若先前沒重構時，是修改全域變數的情況，所以若不可再用:=推斷一次，會導致操作gorm時報錯：invalid memory address or nil pointer dereference
	var err error
	DB, err := gorm.Open(mysql.Open(config), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatalf("Connect to DB err: %v\n", err)
		return DB, err
	}
	fmt.Printf("已讀取mysql資料庫:%v\n", DB)
	return DB, nil
}
