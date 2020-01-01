package cmd

import (
  "github.com/remicaumette/zaap.sh/cmd/zaapctl/cmd/create"
  apiclient "github.com/remicaumette/zaap.sh/pkg/client"
  "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
)

func NewRootCmd(client *apiclient.Client) *cobra.Command {
  cmd := &cobra.Command{
    Use: "zaapctl",
    Short: "zaapctl controls the Zaap server",
    SilenceUsage: true,
    Run: func(cmd *cobra.Command, args []string) {
      if err := cmd.Help(); err != nil {
        logrus.WithError(err).Fatal("failed to show help")
      }
    },
  }
  cmd.AddCommand(create.NewCreateCmd(client))
  return cmd
}
