package controller

import (
	"github.com/chzyer/readline"
	"github.com/jasperstritzke/cubid/cmd/controller/controller_network"
	"github.com/jasperstritzke/cubid/cmd/controller/template"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/commandline"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/security"
	"io"
	"strings"
)

var controllerConfig Config
var controllerServer *controller_network.ControllerServer

func Main() {
	logger.Info("Booting up controller...")

	loadConfig()

	startServer()

	defer InitiateShutdown()

	//initialize key if not existing
	security.InitControllerKey()
	security.LoadControllerKey()

	template.LoadTemplates()

	//Always must be executed last because it's a blocking task.
	startCommandLine()
}

func startServer() *controller_network.ControllerServer {
	server := controller_network.NewControllerServer(controllerConfig.Host)
	controllerServer = server
	go server.Server.Start()
	return server
}

func loadConfig() {
	configPath := "base.json"
	err := config.InitConfigIfNotExists(configPath, config.WrapExistingConfig(DefaultConfig))
	if err != nil {
		panic(err)
	}

	err = config.LoadConfig(configPath, &controllerConfig)
	if err != nil {
		panic(err)
	}
}

func startCommandLine() {
	completer := readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("quit"),
	)

	cmdLine := commandline.NewCommandLine(completer, "Controller")

cmdLoop:
	for {
		line, err := cmdLine.Line.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "quit"):
			InitiateShutdown()
			break cmdLoop
		}
	}
}

var shutdown = false

func InitiateShutdown() {
	if shutdown {
		return
	}
	shutdown = true

	logger.Warn("Shutting down...")

	controllerServer.Server.Stop()

	logger.Info("Successfully stopped all systems. Bye.")
}
