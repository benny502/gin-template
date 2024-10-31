package config

type Log struct {
	Level string `mapstructure:"level" json:"level" yaml:"level"`
	Path  string `mapstructure:"path" json:"path" yaml:"path"`
}
