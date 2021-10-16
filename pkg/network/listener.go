package network

type ConnectionFunc = func(endpoint *Endpoint)
type PacketListenerFunc = func(endpoint *Endpoint, packet Packet)

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
