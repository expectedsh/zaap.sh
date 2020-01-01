package cmd

import (
  "github.com/remicaumette/zaap.sh/cmd/zaapctl/cmd/app"
  "github.com/remicaumette/zaap.sh/pkg/api"
  "github.com/spf13/cobra"
)

func NewRootCmd(client *api.Client) *cobra.Command {
  cmd := &cobra.Command{
    Use: "zaapctl",
    SilenceErrors: true,
    SilenceUsage: true,
  }
  cmd.AddCommand(app.NewCmd(client))
  return cmd
}
