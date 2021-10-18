package executor

import (
	"github.com/chzyer/readline"
	"github.com/jasperstritzke/cubid/cmd/executor/executor_network"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/commandline"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/security"
	"io"
	"os"
	"strings"
)

var controllerConfig Config

var (
	executorServer *executor_network.ExecutorServer = nil
	executorClient *executor_network.ExecutorClient = nil
)

func Main() {
	logger.Info("Booting up executor...")

	loadConfig()

	security.LoadControllerKey()

	go connectClient()
	go startServer()

	defer InitiateShutdown()

	//Always must be executed last because it's a blocking task.
	startCommandLine()
}

func startServer() *executor_network.ExecutorServer {
	server := executor_network.NewExecutorServer(controllerConfig.Host)
	executorServer = server

	server.Server.Start()
	return server
}

func connectClient() *executor_network.ExecutorClient {
	client := executor_network.NewExecutorClient(controllerConfig.ControllerHost)
	executorClient = client

	client.Connect()
	return client
}

func loadConfig() {
	configPath := "base.json"
	//Initialize config and run through form if no config exists
	err := config.InitConfigIfNotExists(configPath, configForm)

	if err != nil {
		logger.Error("Unable to create configuration.")
		os.Exit(1)
	}

	err = config.LoadConfig(configPath, &controllerConfig)

	if err != nil {
		logger.Error("Unable to load configuration")
		os.Exit(1)
	}
}

func startCommandLine() {
	completer := readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("quit"),
	)

	cmdLine := commandline.NewCommandLine(completer, "Executor")

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
			logger.Warn("Shutdown initiated...")
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

	executorClient.Client.Disconnect()
	executorServer.Server.Stop()

	logger.Info("Successfully stopped all systems. Bye.")
}
