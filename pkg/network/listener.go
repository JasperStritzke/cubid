package network

import "github.com/jasperstritzke/cubid/pkg/network/packet"

type ConnectionFunc = func(endpoint *Endpoint)
type PacketListenerFunc = func(endpoint *Endpoint, packet packet.Packet)

type Listener struct {
	ConnectListener    ConnectionFunc
	PacketListener     PacketListenerFunc
	DisconnectListener ConnectionFunc
}

func NewListener(onConnect ConnectionFunc, packetListener PacketListenerFunc, onDisconnect ConnectionFunc) *Listener {
	return &Listener{
		ConnectListener:    onConnect,
		PacketListener:     packetListener,
		DisconnectListener: onDisconnect,
	}
}
