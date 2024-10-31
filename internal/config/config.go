package config

import (
	path "bookmark/internal/utils"
	"fmt"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	rootPath = path.RootPath()
)

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type Config struct {
	v *viper.Viper
}

func (c *Config) Load(configPath string) (*Configuration, error) {

	conf := &Configuration{}

	if !filepath.IsAbs(configPath) {
		configPath = filepath.Join(rootPath, configPath)
	}

	fmt.Println("load config:" + configPath)

	c.v.SetConfigType("yaml")
	c.v.SetConfigFile(configPath)
	if err := c.v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file failed: %s \n", err)
	}

	if err := c.v.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("unmarshal config file failed: %s \n", err)
	}

	c.v.WatchConfig()
	c.v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		defer func() {
			if err := recover(); err != nil {
				log.Println("config file changed, but unmarshal failed:", err)
			}
		}()
		if err := c.v.Unmarshal(conf); err != nil {
			fmt.Println(fmt.Errorf("unmarshal config file failed: %s \n", err))
		}
	})

	return conf, nil
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		v: v,
	}
}
