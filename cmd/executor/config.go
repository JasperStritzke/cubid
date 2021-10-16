package executor

type Config struct {
	ControllerHost string `json:"controllerHost"`
	Host           string `json:"host"`
}

var DefaultConfig = Config{
	ControllerHost: "0.0.0.0:4020",
	Host:           "0.0.0.0:4040",
}
