package logger

import (
	"github.com/ivpusic/golog"
	"github.com/jasperstritzke/cubid/pkg/console/color"
	"github.com/jasperstritzke/cubid/pkg/util/fileutil"
	"time"
)

func printHeader() {
	Logger.Info(
		color.Blue + "\n   _____      _     _     _ " +
			"\n / ____|    | |   (_)   | |" +
			"\n| |    _   _| |__  _  __| |" +
			"\n| |   | | | | '_ \\| |/ _` |" +
			"\n| |___| |_| | |_) | | (_| |" +
			"\n \\_____\\__,_|_.__/|_|\\__,_|" +
			"\n" + color.Reset,
	)
}

var Logger *golog.Logger

func init() {
	Logger = golog.Default

	cubidAppender := &CubidAppender{}
	Logger.Disable("github.com/ivpusic/golog/stdout")
	Logger.Enable(cubidAppender)
	golog.Default = Logger

	printHeader()
}

func ActivateLogs() {
	fileName := time.Now().Format("15-04-05 02-01-2006" + ".log")
	path := "logs/" + fileName

	err := fileutil.CreateIfNotExists(path)

	if err != nil {
		panic(err)
	}

	Logger.Enable(File(golog.Conf{
		"path": path,
	}))
}

//Methods for quick access

func Info(msg string) {
	Logger.Info(msg)
}

func Infof(msg string, params ...interface{}) {
	Logger.Infof(msg, params)
}

func Warn(msg string) {
	Logger.Warn(msg)
}

func Warnf(msg string, params ...interface{}) {
	Logger.Warnf(msg, params)
}

func Error(msg string) {
	Logger.Error(msg)
}

func Errorf(msg string, params ...interface{}) {
	Logger.Errorf(msg, params)
}

func Debug(msg string) {
	Logger.Debug(msg)
}

func Debugf(msg string, params ...interface{}) {
	Logger.Debugf(msg, params)
}
