package template

import (
	"fmt"
	"github.com/jasperstritzke/cubid/pkg/console/color"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/model"
	"github.com/jasperstritzke/cubid/pkg/util/timeutil"
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
	if len(line) == 8 {
		templateCommandHelp()
		return
	}

	line = line[9:]

	if strings.ToLower(line) == "list" {
		logger.Info("Loaded templates: ")
		for key, value := range templates {
			logger.Info("• " + key + ":")

			for _, template := range value {
				logger.Info("  Name: " + template.Name)

				logger.Info("  • Type: " + color.Green + template.ProxyAsString() + color.Reset)
				logger.Info("  • Version: " + template.Version.Display)
				logger.Info("")
			}
		}
		return
	}

	if strings.ToLower(line) == "versions" {
		logger.Info("Available pre-configured versions:")

		for _, version := range model.VersionsAsArray {
			logger.Info(version.Display)

			logger.Info("  • Type: " + version.ProxyAsString())
			if len(version.Mention) > 0 {
				logger.Info("  • Mention: " + version.Mention)
			}
			logger.Info("  • BuildURL: " + version.BuildURL)
		}

		return
	}

	if strings.ToLower(line) == "reload" {
		logger.Warn("Reloading all templates...")

		stopWatch := timeutil.StopWatch{}
		stopWatch.Start()

		LoadTemplates()

		stopWatch.Stop()

		logger.Info("Successfully reloaded all templates in " + fmt.Sprint(stopWatch.GetDurationInMilliseconds()) + "ms.")
		return
	}

	logger.Warn("Unknown command. Use 'template help' for help.")
}
