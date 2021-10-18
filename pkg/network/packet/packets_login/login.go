package packets

import "github.com/jasperstritzke/cubid/pkg/model"

const LoginRequestAction = "REQUEST_LOGIN"

type PacketAuthLoginRequest struct {
	Info model.ExecutorInfo `json:"info"`
	Key  string             `json:"key"`
}

const LoginResponseAction = "RESPONSE_LOGIN"

type PacketAuthLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
