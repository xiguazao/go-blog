package ansible_service

import (
	"context"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

type ConnectOptions struct {

	// smart
	Connection string

	// root
	User string

	//"10.100.2.118,10.100.2.164"
	Inventory string

	// {"ansible-runner/site.yml", "ansible-runner/site2.yml"}
	Playbooks []string
}

func (con_opt *ConnectOptions) PlayBookExecutor() (bool, error) {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: con_opt.Connection,
		User:       con_opt.User,
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: con_opt.Inventory,
	}

	ansiblePlaybook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         con_opt.Playbooks,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := ansiblePlaybook.Run(context.TODO())
	if err != nil {
		//panic(err)
		return false, err
	}
	return true, nil
}
