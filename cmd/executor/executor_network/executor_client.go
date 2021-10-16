package executor_network

import (
	"github.com/jasperstritzke/cubid/pkg/network"
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

func onConnectToServer(endpoint *network.Endpoint) {
}

func onDisconnectFromServer(endpoint *network.Endpoint) {
}

func onPacketFromServer(endpoint *network.Endpoint, packet network.Packet) {
}
