package data

import (
	"bookmark/internal/config"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	db     *gorm.DB
	conf   *config.Configuration
	writer io.Writer
}

func NewData(conf *config.Configuration, writer io.Writer) (*Data, func(), error) {

	data := &Data{
		conf:   conf,
		writer: writer,
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

	return logger.New(log.New(d.writer, "\r\n", log.LstdFlags), logger.Config{
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

var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewClassRepo, NewItemRepo)
