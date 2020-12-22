package config

type DBconf struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Charset  string //編碼
}
