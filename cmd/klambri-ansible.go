package main

import (
	"fmt"
	"klambri-ansible/internal/ansible"
	"klambri-ansible/internal/models"
	"log"
)

func main() {
	playbookConfig := &models.PlaybookConfig{
		PlaybookName: "../playbooks/hello-world.yaml",
		Hosts:        []string{"127.0.0.1", "127.0.0.1"},
	}

	err := ansible.Run(playbookConfig)
	if err != nil {
		log.Fatalf("Error running playbook: %v", err)
	} else {
		fmt.Println("Playbook executed successfully!")
	}
}
