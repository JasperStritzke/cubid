package main

import (
	"flag"
	"github.com/jasperstritzke/cubid/cmd/controller"
	"github.com/jasperstritzke/cubid/cmd/executor"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	_ "github.com/jasperstritzke/cubid/pkg/console/logger"
)

func main() {
	logger.Warn("Loading...")

	startup := flag.String("startup", "none", "specify whether you cau want to startup a CONTROLLER or EXECUTOR")
	noLogs := flag.Bool("nologs", false, "specify whether noLogs should be written.")

	flag.Parse()

	if !*noLogs {
		logger.ActiveLogs()
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
