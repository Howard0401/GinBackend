package settings

type SMTP struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	SSL      bool   `mapstructure:"ssl" yaml:"ssl"`
	UserName string `mapstructure:"user_name" yaml:"user_name"`
	Password string `mapsturcture:"password" yaml:"password"`
	From     string `mapstructure:"from" yaml:"from"`
	To       string `mapstructure:"to" yaml:"to"`
}
