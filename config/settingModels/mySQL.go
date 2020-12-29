package settings

type MySQL struct {
	Name     string `maptructure:"name" yaml:"name"`
	Host     string `maptructure:"host" yaml:"host"`
	UserName string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}
