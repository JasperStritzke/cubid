package controller_network

import (
	"github.com/jasperstritzke/cubid/pkg/network"
	"github.com/jasperstritzke/cubid/pkg/network/packet"
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

func onPacket(endpoint *network.Endpoint, packet packet.Packet) {
}
