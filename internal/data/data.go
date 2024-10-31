package data

import (
	"bookmark/internal/config"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	path "bookmark/internal/utils"
)

type Data struct {
	db   *gorm.DB
	conf *config.Configuration
}

func NewData(conf *config.Configuration) (*Data, func(), error) {

	data := &Data{
		conf: conf,
	}
	err := data.Connect()
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
	}
	return data, cleanup, nil

}

func (d *Data) Connect() error {
	switch d.conf.Database.Driver {
	case "mysql":
		return d.mysql()
	default:
		return fmt.Errorf("not support database %s", d.conf.Database.Driver)
	}
}

func (d *Data) logger() logger.Interface {
	logFileDir := d.conf.Log.Path
	if !filepath.IsAbs(logFileDir) {
		logFileDir = filepath.Join(path.RootPath(), logFileDir)
	}

	if ok, _ := path.Exists(logFileDir); !ok {
		os.Mkdir(logFileDir, os.ModePerm)
	}

	now := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s/logs/%s.log", path.RootPath(), now)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
	}

	var logLevel logger.LogLevel
	switch d.conf.Log.Level {
	case "debug":
		logLevel = logger.Info
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	}

	return logger.New(log.New(io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      logLevel,
		Colorful:      true,
	})
}

func (d *Data) mysql() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.conf.Database.User, d.conf.Database.Password, d.conf.Database.Host, d.conf.Database.Port, d.conf.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: d.logger()})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

var ProviderSet = wire.NewSet(NewData, NewUserRepo)
