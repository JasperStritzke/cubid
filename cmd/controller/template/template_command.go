package template

import (
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"strings"
)

func templateCommandHelp() {
	logger.Info("template help - this help page")
	logger.Info("template create <Name> <Proxy (Yes/No) <Version/ 'none'> - create a template")
	logger.Info("template createEmpty <Name> <Proxy (Yes/No> - create an empty template")
	logger.Info("template versions - list all possible versions>")
	logger.Info("template reload - reload all templates from disk")
	logger.Info("template list - lists all templates")
}

func HandleTemplateCommand(line string) {
	//template X

	if strings.ToLower(line) == "template" {
		templateCommandHelp()
		return
	}

	logger.Warn("Unknown command. Use 'template help' for help.")
}
