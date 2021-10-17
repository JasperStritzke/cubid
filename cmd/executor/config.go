package executor

import (
	"fmt"
	"github.com/jasperstritzke/cubid/pkg/console"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/model"
	"os"
)

type Config struct {
	ControllerHost string             `json:"controllerHost"`
	Host           string             `json:"host"`
	Info           model.ExecutorInfo `json:"info"`
}

var DefaultConfig = Config{
	ControllerHost: "0.0.0.0:4020",
	Host:           "0.0.0.0:4040",
	Info: model.ExecutorInfo{
		Name:      "",
		MaxMemory: 2048,
		MaxCPU:    80,
	},
}

func configForm() interface{} {
	var name string
	console.AskFor(
		"Since you're running this executor the first time, please define a name.",
		"Name: ", &name,
	)

	if len(name) == 0 {
		logger.Error("The name must consist of at least 1 character.")
		os.Exit(1)
	}

	logger.Info("Name \"" + name + "\" accepted.")

	DefaultConfig.Info.Name = name

	var maxMemory int
	console.AskFor(
		"Please define the maximum amount of memory this executor is allowed to start services with.",
		"Max-Memory (MB): ", &maxMemory,
	)

	if maxMemory < 256 {
		logger.Error("You must provide each executor with at least 256MB of memory.")
		os.Exit(1)
	}

	if maxMemory > 1024*1024 {
		logger.Error("You can't provide more than 1TB of memory.")
		os.Exit(1)
	}

	DefaultConfig.Info.MaxMemory = maxMemory
	logger.Info("Max-Memory \"" + fmt.Sprint(maxMemory) + "MB\" accepted.")

	var maxCPUUsage int
	console.AskFor(
		"Please define the maximum CPU-Usage which shouldn't be crossed. (do not add the % char)",
		"Max-CPU-Usage (in %): ",
		&maxCPUUsage,
	)

	if maxCPUUsage < 1 || maxCPUUsage > 100 {
		logger.Error("The maximum CPU Usage must be a number between 1 and 100.")
	}

	DefaultConfig.Info.MaxCPU = maxCPUUsage
	logger.Info("Max-CPU-Usage \"" + fmt.Sprint(maxCPUUsage) + "%\" accepted.")

	fmt.Println()
	logger.Info("Setup completed.")
	fmt.Println()

	return DefaultConfig
}
