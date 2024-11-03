package main

import (
	"bookmark/internal/config"
	"bookmark/internal/middleware/cache"
	"flag"

	"github.com/spf13/viper"
)

var (
	flagconf string
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
	cache := cache.NewCache(cache.NewNoCache())
	app, cleanup, err := wireApp(conf, cache)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
