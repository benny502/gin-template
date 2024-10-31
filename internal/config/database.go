package config

type Database struct {
	Driver   string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
}
