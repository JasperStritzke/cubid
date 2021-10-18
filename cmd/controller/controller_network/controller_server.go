package controller_network

import (
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/network"
	"github.com/jasperstritzke/cubid/pkg/network/packet"
	"github.com/jasperstritzke/cubid/pkg/network/packet/packets_login"
	"github.com/jasperstritzke/cubid/pkg/security"
)

type ControllerServer struct {
	Server *network.Server
}

func NewControllerServer(host string) *ControllerServer {
	server := network.NewServer(
		host,
		*network.NewListener(
			onConnect, onPacket, onDisconnect,
		),
	)

	return &ControllerServer{
		Server: server,
	}
}

func onConnect(endpoint *network.Endpoint) {
}

func onDisconnect(endpoint *network.Endpoint) {
}

func onPacket(endpoint *network.Endpoint, pkg packet.Packet) {
	switch pkg.Action {
	case packets.LoginRequestAction:
		var request packets.PacketAuthLoginRequest
		err := pkg.GetPayload(&request)

		if err != nil {
			logger.Warn("Received login request but was unable to resolve request. Request rejected and connection closed.")
			_ = endpoint.Close()
			return
		}

		logger.Warn("Executor " + request.Info.Name + " [" + endpoint.IP() + "] logging in. Validating Key...")

		if !security.IsHashedKeyValid(request.Key) {
			logger.Error("Executor " + request.Info.Name + " was rejected due to invalid key.")

			response, _ := packet.ByPayload(packets.LoginResponseAction,
				packets.PacketAuthLoginResponse{
					Success: false,
					Message: "invalid key",
				})

			err = endpoint.SendPacket(response)

			if err != nil {
				_ = endpoint.Close()
				return
			}
			return
		}

		logger.Info("Executor " + request.Info.Name + " logged in.")

		response, _ := packet.ByPayload(packets.LoginResponseAction, packets.PacketAuthLoginResponse{
			Success: true,
			Message: "logged in",
		})

		err = endpoint.SendPacket(response)
		if err != nil {
			_ = endpoint.Close()
			return
		}

		break
	}
}
