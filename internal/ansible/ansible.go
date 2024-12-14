package ansible

import "klambri-ansible/internal/models"

func Run(playbookConfig *models.PlaybookConfig) error {
	inventory := createInventory(playbookConfig.Hosts)

	// if err := runPlaybook(playbookConfig.PlaybookName, inventory, map[string]interface{}{
	// 	"package": "python3-venv",
	// })
	err := runPlaybook(playbookConfig.PlaybookName, inventory, map[string]interface{}{
		"package": "firefox",
	})

	if err != nil {
		return err
	}

	// if err := runPlaybook(playbookConfig.PlaybookName, inventory, nil); err != nil {
	// 	return err
	// }

	return nil
}
