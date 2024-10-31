package config

type App struct {
	Name    string `mapstructure:"env" json:"env" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
}