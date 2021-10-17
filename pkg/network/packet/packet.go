package packet

import (
	"encoding/json"
)

type Packet struct {
	Action  string  `json:"action"`
	Payload Payload `json:"payload"`
}

type Payload = string

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
