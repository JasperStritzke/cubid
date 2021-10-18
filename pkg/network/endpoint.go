package network

import (
	"encoding/json"
	"github.com/jasperstritzke/cubid/pkg/network/packet"
	"net"
)

type Endpoint struct {
	conn    net.Conn
	encoder *json.Encoder
}

func (endpoint *Endpoint) SendPacket(packet *packet.Packet) error {
	err := endpoint.encoder.Encode(packet)
	return err
}

func (endpoint *Endpoint) SendPackets(packets ...*packet.Packet) error {
	for _, p := range packets {
		err := endpoint.SendPacket(p)

		if err != nil {
			return err
		}
	}

	return nil
}

func (endpoint *Endpoint) Close() error {
	return endpoint.conn.Close()
}

func (endpoint *Endpoint) IP() string {
	return endpoint.conn.RemoteAddr().String()
}
