package logger

import (
	"fmt"
	"github.com/ivpusic/golog"
	"github.com/jasperstritzke/cubid/pkg/console/color"
	"os"
)

type FileAppender struct {
	path string
}

func (fa *FileAppender) Id() string {
	return "file_writer"
}

func (fa *FileAppender) Append(log golog.Log) {
	f, err := os.OpenFile(fa.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil || f == nil {
		fmt.Println(err)
		if log.Logger.DoPanic {
			panic(err)
		}
	}

	defer f.Close()

	line := []byte(
		color.EraseColor(
			log.Time.Format("02.01.2006") + " " + log.Level.Name + " " + log.Message,
		),
	)

	if err != nil {
		fmt.Println(err.Error())
		if log.Logger.DoPanic {
			panic(err)
		}
	}

	line = append(line, byte('\n'))

	f.Write(line)
	f.Sync()
}

func File(cnf golog.Conf) *FileAppender {
	return &FileAppender{
		path: cnf["path"],
	}
}
