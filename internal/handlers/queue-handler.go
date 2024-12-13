package handlers

import (
	"encoding/json"
	"klambri-ansible/internal/ansible"
	"klambri-ansible/internal/models"
	"log"

	"github.com/nats-io/nats.go"
)

func OnQueueMessageReceived(msg *nats.Msg) {
	var playbook models.PlaybookConfig
	err := json.Unmarshal(msg.Data, &playbook)
	if err != nil {
		log.Printf("Error unmarshalling message: %s\n", err)
		return
	}

	err = ansible.Run(&playbook)
	if err != nil {
		log.Printf("Error running Ansible: %s\n", err)
		return
	}
}
