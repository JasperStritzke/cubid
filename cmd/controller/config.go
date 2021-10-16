package controller

type Config struct {
	Host string `json:"host"`
}

var DefaultConfig = Config{
	Host: "0.0.0.0:4020",
}
