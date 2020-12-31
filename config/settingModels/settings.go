package settings

type ConfigSettings struct {
	System System `mapsturcture:"system" yaml:"system"`
	MySQL  MySQL  `mapstructure:"mysql" yaml:"mysql"`
	JWT    JWT    `mapstructure:"jwt" yaml:"jwt"`
	Log    Log    `mapstructure:"log" yaml:"log"`
	SMTP   SMTP   `mpatructure:"smtp" yaml:"smtp"`
}
