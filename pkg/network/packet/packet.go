package packet

import (
	"encoding/json"
	"errors"
)

type Packet struct {
	Action  string  `json:"action"`
	Payload Payload `json:"payload"`
}

type Payload = string

func ByPayload(action string, v interface{}) (*Packet, error) {
	if len(action) == 0 {
		return nil, errors.New("action must be at least 1 character")
	}

	packet := Packet{
		Action: action,
	}

	err := packet.SetPayLoad(v)

	return &packet, err
}

func (packet *Packet) SetPayLoad(v interface{}) error {
	bytes, err := json.Marshal(v)
	packet.Payload = string(bytes)

	return err
}

func (packet *Packet) GetPayload(v interface{}) error {
	return json.Unmarshal(
		[]byte(packet.Payload), v,
	)
}
