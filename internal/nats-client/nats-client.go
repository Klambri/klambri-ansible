package natsclient

import (
	"github.com/nats-io/nats.go"
)

func Run() error {
	nc, err := connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	subscribe(nc, "playbookQueue")

	select {}
}
