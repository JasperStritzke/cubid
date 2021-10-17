package packets

import "github.com/jasperstritzke/cubid/pkg/model"

type PacketAuthLoginRequest struct {
	Info model.ExecutorInfo `json:"info"`
	Key  string             `json:"key"`
}

type PacketAuthLoginResponse struct {
	Success bool `json:"success"`
}
