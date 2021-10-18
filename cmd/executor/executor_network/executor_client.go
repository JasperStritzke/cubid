package executor_network

import (
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/model"
	"github.com/jasperstritzke/cubid/pkg/network"
	"github.com/jasperstritzke/cubid/pkg/network/packet"
	"github.com/jasperstritzke/cubid/pkg/network/packet/packets_login"
	"github.com/jasperstritzke/cubid/pkg/security"
	"os"
)

type ExecutorClient struct {
	Client *network.Client
}

func NewExecutorClient(host string) *ExecutorClient {
	client := network.NewClient(
		host,
		*network.NewListener(
			onConnectToServer, onPacketFromServer, onDisconnectFromServer,
		),
	)

	return &ExecutorClient{
		Client: client,
	}
}

func (c *ExecutorClient) Connect() {
	c.Client.Connect()
}

func onConnectToServer(endpoint *network.Endpoint) {
	loginRequest, err := packet.ByPayload(packets.LoginRequestAction, &packets.PacketAuthLoginRequest{
		Info: model.ExecutorInfo{},
		Key:  security.GetHashedKey(),
	})

	if err != nil {
		logger.Error("Unable to request login: " + err.Error())
		os.Exit(1)
	}

	err = endpoint.SendPacket(loginRequest)

	if err != nil {
		logger.Error("Unable to request login: " + err.Error())
		os.Exit(1)
	}
}

func onDisconnectFromServer(endpoint *network.Endpoint) {
}

func onPacketFromServer(endpoint *network.Endpoint, pkg packet.Packet) {
	switch pkg.Action {
	case packets.LoginRequestAction:
		var response packets.PacketAuthLoginResponse
		_ = pkg.GetPayload(&response)

		if !response.Success {
			logger.Error("Login to controller was rejected: " + response.Message + ".")
			os.Exit(1)
		}

		logger.Info("Login to controller was accepted.")

		break
	}
}
