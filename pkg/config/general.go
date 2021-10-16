package config

type GeneralConfig struct {
	Type string
}

func (config *GeneralConfig) IsController() bool {
	return config.Type == TypeController
}

func (config *GeneralConfig) IsExecutor() bool {
	return config.Type == TypeExecutor
}

const (
	TypeController = "CONTROLLER"
	TypeExecutor   = "EXECUTOR"
)
