package log

import (
	"bookmark/internal/config"
	path "bookmark/internal/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Writer struct {
	conf *config.Configuration
}

func (w *Writer) Write(p []byte) (n int, err error) {
	logFileDir := w.conf.Log.Path
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
		return 0, err
	}
	defer f.Close()
	return f.Write(p)
}

func NewWriter(conf *config.Configuration) io.Writer {

	return io.MultiWriter(os.Stdout, &Writer{
		conf: conf,
	})
}
