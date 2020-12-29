package settings

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	FileName   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
}
