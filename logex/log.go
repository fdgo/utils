package logex

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

func init()  {
	fileName := "log"
	log.AddHook(newLfsHook(fileName))
}

func newLfsHook(fileName string) log.Hook {
	writer, err := rotatelogs.New(
		fileName+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(3)), //日志分割周期
		//rotatelogs.WithRotationCount(8),                      //设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*24), //清理前最多保存时间
	)
	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}
	log.SetLevel(log.DebugLevel)
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{})

	return lfsHook
}
