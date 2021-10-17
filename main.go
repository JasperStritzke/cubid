package main

import (
	"flag"
	"fmt"
	"github.com/jasperstritzke/cubid/cmd/controller"
	"github.com/jasperstritzke/cubid/cmd/executor"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/color"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	_ "github.com/jasperstritzke/cubid/pkg/console/logger"
	"os"
)

const version = "0.1-ALPHA"

func main() {
	logger.Info("Starting cubid v" + version + " by Jasper Stritzke")

	startup := flag.String("startup", "none", "specify whether you cau want to startup a CONTROLLER or EXECUTOR")
	noLogs := flag.Bool("nologs", false, "specify whether noLogs should be written.")
	help := flag.Bool("help", false, "shows all possible flags.")
	reset := flag.Bool("reset", false, "resets all configurations and logs on startup.")

	flag.Parse()

	if *help {
		optional := color.Green + "(Optional)" + color.Reset
		required := color.Red + "(Required)" + color.Reset

		logger.Info("All possible flags are:")
		logger.Info(optional + " -help | displays all possible flags")
		logger.Warn(required + " -startup=<CONTROLLER|EXECUTOR> | decide whether you want to start a controller or executor")
		logger.Info(optional + " -nologs | disable logging to files. This is useful when you have to restart the cloud a lot (in development for example)")
		logger.Info(optional + " -reset | resets all configurations and logs on startup.")
		fmt.Println(color.Yellow + "(!) Please note, that the software won't start if you haven't provided all required flags." + color.Reset)
		return
	}

	if *reset {
		logger.Warn("Deleting configurations and logs...")
		_ = os.RemoveAll("config/")
		_ = os.RemoveAll("logs/")
	}

	if !*noLogs {
		logger.ActivateLogs()
	} else {
		logger.Warn("This service is running with logging deactivated. This is not recommended in production environment.")
	}

	switch *startup {
	case config.TypeController:
		controller.Main()
		break
	case config.TypeExecutor:
		executor.Main()
		break
	default:
		logger.Debug("Invalid or undefined startup option.")
		panic("Invalid or undefined startup option.")
		return
	}
}
