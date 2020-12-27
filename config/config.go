package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Vp   *viper.Viper
}

//Constructor
func InitViper(name string) (Config, error) {
	c := Config{
		Name: name,
		Vp:   viper.New(),
	}
	if err := c.loadConfig(); err != nil {
		return c, err
	}
	c.watchConfig()
	return c, nil
}

//產生viper的log
func (c *Config) loadConfig() error {
	if c.Name != "" {
		c.Vp.SetConfigFile(c.Name)
	} else {
		c.Vp.AddConfigPath("config")
		c.Vp.SetConfigName("config")
	}
	c.Vp.SetConfigType("yaml")
	if err := c.Vp.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

//即時監聽Config的變化，記入log中
func (c *Config) watchConfig() {
	c.Vp.WatchConfig()
	c.Vp.OnConfigChange(func(in fsnotify.Event) {
		//這邊如果在生產環境，可以用channel配合goroutine紀錄(因為可能會有非同步的情形)
		log.Printf("Config file changed: %s", in.Name)
	})
}
