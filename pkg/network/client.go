package network

import (
	"encoding/json"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	packet2 "github.com/jasperstritzke/cubid/pkg/network/packet"
	"net"
)

type Client struct {
	service   string
	conn      net.Conn
	listeners Listener
	Endpoint  *Endpoint
}

func NewClient(service string, listeners Listener) *Client {
	return &Client{
		service:   service,
		listeners: listeners,
	}
}

func (client *Client) Disconnect() {
	_ = client.conn.Close()
}

//Connect is recommended to be executed asynchronously
func (client *Client) Connect() {
	logger.Info("Connecting...")

	clientConnection, err := net.Dial("tcp", client.service)
	client.conn = clientConnection

	if err != nil {
		panic(err)
	}

	logger.Info("Connected to " + client.service + ".")

	encoder := json.NewEncoder(clientConnection)
	decoder := json.NewDecoder(clientConnection)

	endpoint := Endpoint{
		conn:    clientConnection,
		encoder: encoder,
	}

	client.Endpoint = &endpoint

	client.listeners.ConnectListener(&endpoint)

	for {
		var packet packet2.Packet
		decodeErr := decoder.Decode(&packet)

		if decodeErr != nil {
			_ = clientConnection.Close()
			break
		}

		client.listeners.PacketListener(&endpoint, packet)
	}

	client.listeners.DisconnectListener(&endpoint)
	logger.Info("Connection closed.")
}
