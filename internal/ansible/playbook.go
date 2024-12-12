package ansible

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func runPlaybook(playbookName string, inventory string) error {
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  inventory,
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbookName),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
	)

	err := exec.Execute(context.Background())
	if err != nil {
		return err
	}

	return nil
}
