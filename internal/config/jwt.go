package config

type Jwt struct {
	JwtKey string `mapstructure:"jwt_key" json:"jwt_key" yaml:"jwt_key"`
	Issuer string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	Expttl int    `mapstructure:"expttl" json:"expttl" yaml:"expttl"`
}
