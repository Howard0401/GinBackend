package settings

import "time"

type System struct {
	Name          string        `mapstructue:"name" yaml:"name"`
	FirstDB       string        `mapstructure:"first_db" yaml:"first_db"`
	Mode          string        `mapstructure:"mode" yaml:"mode"`
	Port          string        `mapstructure:"port" yaml:"port"`
	Url           string        `mapstructure:"url" yaml:"url"`
	MaxCheckCount int           `mapstructure:"max_check_count" yaml:"max_check_count"`
	Timeout       time.Duration `mapstructure:"timeout" yaml:"timeout"`
}
