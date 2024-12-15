package main

import (
	natsclient "klambri-ansible/internal/nats-client"
	"log"
)

func main() {
	err := natsclient.Run()
	if err != nil {
		log.Fatalf("Error running NATS client: %s", err)
	}
}
