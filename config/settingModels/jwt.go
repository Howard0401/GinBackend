package settings

type JWT struct {
	Secret string `mapstructure:"secret" yaml:"secret"`
	Issuer string `mapstructure:"issuer" yaml:"issuer"`
	Expire int    `mapstructure:"expire" yaml:"expire"`
}
