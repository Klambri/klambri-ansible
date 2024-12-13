package ansible

import "klambri-ansible/internal/models"

func Run(playbookConfig *models.PlaybookConfig) error {
	inventory := createInventory(playbookConfig.Hosts)

	if err := runPlaybook(playbookConfig.PlaybookName, inventory); err != nil {
		return err
	}

	return nil
}
