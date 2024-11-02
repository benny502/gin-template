package log

import (
	"bookmark/internal/config"
	path "bookmark/internal/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
}

type logrus_ struct {
	*logrus.Logger
}

func NewLogger(conf *config.Configuration) (Logger, error) {
	logger := logrus.New()
	switch conf.Log.Level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	logFileDir := conf.Log.Path
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
		return nil, err
	}
	logger.SetOutput(io.MultiWriter(os.Stdout, f))
	return &logrus_{logger}, nil
}
