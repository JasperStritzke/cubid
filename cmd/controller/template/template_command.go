package template

import (
	"fmt"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/model"
	"github.com/jasperstritzke/cubid/pkg/util/timeutil"
	"strings"
)

func showTemplateCommandHelp() {
	logger.Info("template - this help page")

	logger.Info("template create <Group> <Name> <Version> - create a template")
	logger.Info("template createEmpty <Group> <Name> - create an empty template")

	logger.Info("template disable <Group> <Name> - disable a template")
	logger.Info("template enable <Group> <Name> - enable a template")

	logger.Info("template versions - list all possible versions>")
	logger.Info("template reload - reload all templates from disk")
	logger.Info("template list - lists all templates")
}

func HandleTemplateCommand(line string) {
	if len(line) == 8 {
		showTemplateCommandHelp()
		return
	}

	line = line[9:]

	subCommandLower := strings.ToLower(line)
	if subCommandLower == "list" {
		logger.Info("Loaded templates: ")
		for key, value := range templates {
			logger.Info("• " + key + ":")

			for _, template := range value {
				logger.Info("  » " + template.Name + " - Version: " + template.Version.Display)
			}

			logger.Info("")
		}
		return
	}

	if subCommandLower == "versions" {
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

	if subCommandLower == "reload" {
		logger.Warn("Reloading all templates...")

		stopWatch := timeutil.StopWatch{}
		stopWatch.Start()

		LoadTemplates()

		stopWatch.Stop()

		logger.Info("Successfully reloaded all templates in " + fmt.Sprint(stopWatch.GetDurationInMilliseconds()) + "ms.")
		return
	}

	if strings.HasPrefix(subCommandLower, "create") {
		if len(subCommandLower) <= 6 {
			logger.Warn("Invalid usage of template create.")
			return
		}

		createSubCommand := line[7:]
		createSubCommands := strings.Split(createSubCommand, " ")

		if len(createSubCommands) < 3 {
			logger.Warn("Invalid usage of template create.")
			return
		}

		group := createSubCommands[0]
		name := createSubCommands[1]
		versionRaw := strings.ToLower(createSubCommands[2])
		logger.Error(versionRaw)
		version, err := model.GetVersionByString(versionRaw)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		err = CreateTemplate(group, name, version)
		if err != nil {
			logger.Error("Error creating template: " + err.Error())
			return
		}

		logger.Info("Template " + group + "/" + name + " successfully created.")
		return
	}

	if strings.HasPrefix(subCommandLower, "disable") || strings.HasPrefix(subCommandLower, "enable") {
		isEnable := strings.HasPrefix(subCommandLower, "enable")

		if len(line) <= 7 {
			logger.Warn("Invalid usage of template enable or disable.")
			return
		}

		var disable string
		if strings.HasPrefix(subCommandLower, "disable") {
			disable = line[8:]
		} else {
			disable = line[7:]
		}

		enableSubCommands := strings.Split(disable, " ")

		if len(enableSubCommands) < 2 {
			logger.Warn("Invalid usage of template enable or disable.")
			return
		}

		group := enableSubCommands[0]
		template := enableSubCommands[1]

		success, err := SetEnabled(group, template, isEnable)

		if err != nil {
			logger.Error("Error: " + err.Error())
			return
		}

		if !success {
			logger.Warn("Template not found!")
		}

		logger.Info("Changes saved. Use template reload to apply changes")
		return
	}

	logger.Warn("Unknown command. Use 'template help' for help.")
	showTemplateCommandHelp()
}
