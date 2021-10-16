package network

import (
	"encoding/json"
	"net"
)

type Endpoint struct {
	conn    net.Conn
	encoder *json.Encoder
}

func (endpoint *Endpoint) SendPacket(packet *Packet) error {
	err := endpoint.encoder.Encode(packet)
	return err
}

func (endpoint *Endpoint) SendPackets(packets ...*Packet) error {
	for _, packet := range packets {
		err := endpoint.SendPacket(packet)

		if err != nil {
			return err
		}
	}

	return nil
}
