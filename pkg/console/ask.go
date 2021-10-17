package console

import (
	"fmt"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
)

func AskFor(question, prompt string, v interface{}) {
	logger.Info(question)
	fmt.Print(prompt)

	_, err := fmt.Scan(v)

	if err != nil {
		return
	}
}
