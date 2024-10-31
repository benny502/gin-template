package main

import (
	"bookmark/internal/config"
	"flag"

	"github.com/spf13/viper"
)

var (
	flagconf string
	flagport string
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs/config.yaml", "config path, eg: -conf config.yaml")
}
func main() {
	flag.Parse()
	config := config.NewConfig(viper.New())
	conf, err := config.Load(flagconf)
	if err != nil {
		panic(err)
	}
	app, cleanup, err := wireApp(conf)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Run(flagconf) // listen and serve on 0.0.0.0:8080
}
