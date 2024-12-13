package natsclient

import (
	"klambri-ansible/internal/handlers"
	"log"

	"github.com/nats-io/nats.go"
)

func subscribe(nc *nats.Conn, subject string) {
	nc.Subscribe(subject, func(msg *nats.Msg) {
		handlers.OnQueueMessageReceived(msg)
	})

	log.Println("Subscribed to:", subject)
}
