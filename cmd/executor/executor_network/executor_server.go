package executor_network

import (
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/network"
)

type ExecutorServer struct {
	Server *network.Server
}

func NewExecutorServer(host string) *ExecutorServer {
	server := network.NewServer(
		host,
		*network.NewListener(
			onClientConnect, onPacketFromClient, onClientDisconnect,
		),
	)

	return &ExecutorServer{
		Server: server,
	}
}

func onClientConnect(endpoint *network.Endpoint) {
	logger.Info("Connecting")
}

func onClientDisconnect(endpoint *network.Endpoint) {
	logger.Info("Disconnecting")
}

func onPacketFromClient(endpoint *network.Endpoint, packet network.Packet) {
}
