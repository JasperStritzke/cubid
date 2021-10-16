package logger

import (
	"fmt"
	"github.com/ivpusic/golog"
	"github.com/jasperstritzke/cubid/pkg/console/color"
)

type CubidAppender struct{}

func getLevelColor(level golog.Level) string {
	switch level.Value {
	case golog.INFO.Value:
		return color.Green
	case golog.WARN.Value:
		return color.Yellow
	case golog.ERROR.Value:
		return color.Red
	default:
		return color.Blue
	}
}

const timeFormat = "15:04:05"

func (s *CubidAppender) Id() string {
	return "cubid"
}

func (s *CubidAppender) Append(log golog.Log) {
	msg := fmt.Sprintf(
		color.Gray+"[%s%s "+color.White+"%s"+color.Gray+"]"+color.Reset+" %s",
		getLevelColor(log.Level),
		log.Level.Name,
		log.Time.Format(timeFormat),
		log.Message,
	)

	fmt.Println(msg)
}
