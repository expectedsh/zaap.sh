package app

import (
  "github.com/remicaumette/zaap.sh/pkg/api"
  "github.com/spf13/cobra"
)

func NewCmd(client *api.Client) *cobra.Command {
  cmd := &cobra.Command{
    Use: "app",
    Short: "Interact with apps",
    Long: `This command groups subcommands for interacting with apps.

  Create an app:

      $ zaapctl app create <name>

  Please see the individual subcommand help for detailed usage information.`,
  }
  cmd.AddCommand(NewCreateCmd(client))
  return cmd
}
