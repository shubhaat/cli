package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flags"
)

type SetHealthCheckCommand struct {
	RequiredArgs flags.SetHealthCheckArgs `positional-args:"yes"`
	usage        interface{}              `usage:"CF_NAME set-health-check APP_NAME ('port' | 'none')"`
}

func (_ SetHealthCheckCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (_ SetHealthCheckCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}