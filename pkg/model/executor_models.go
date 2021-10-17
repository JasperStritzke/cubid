package model

type ExecutorInfo struct {
	Name      string `json:"name"`
	MaxMemory int    `json:"maxMemory"`
	MaxCPU    int    `json:"maxCpu"`
}
