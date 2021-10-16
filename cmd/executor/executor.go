package executor

import (
	"github.com/chzyer/readline"
	"github.com/jasperstritzke/cubid/cmd/executor/executor_network"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/commandline"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"io"
	"strings"
)

var controllerConfig Config

var (
	executorServer *executor_network.ExecutorServer
	executorClient *executor_network.ExecutorClient
)

func Main() {
	logger.Info("Booting up executor...")

	loadConfig()

	connectClient()
	startServer()

	defer InitiateShutdown()

	//Always must be executed last because it's a blocking task.
	startCommandLine()
}

func startServer() *executor_network.ExecutorServer {
	server := executor_network.NewExecutorServer(controllerConfig.Host)
	executorServer = server

	go server.Server.Start()
	return server
}

func connectClient() *executor_network.ExecutorClient {
	client := executor_network.NewExecutorClient(controllerConfig.ControllerHost)
	executorClient = client

	go client.Client.Connect()
	return client
}

func loadConfig() {
	configPath := "config/base.json"
	err := config.InitConfigIfNotExists(configPath, DefaultConfig)
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
