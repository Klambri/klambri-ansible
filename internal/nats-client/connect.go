package natsclient

import (
	"log"

	"github.com/nats-io/nats.go"
)

func connect(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
		return nil, err
	}
	return nc, nil
}
