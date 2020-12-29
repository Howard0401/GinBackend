package settings

type System struct {
	FirstDB       string `mapstructure:"first_db" yaml:"first_db"`
	Mode          string `mapstructure:"mode" yaml:"mode"`
	Port          string `mapstructure:"port" yaml:"port"`
	Url           string `mapstructure:"url" yaml:"url"`
	MaxCheckCount int    `mapstructure:"max_check_count" yaml:"max_check_count"`
}
