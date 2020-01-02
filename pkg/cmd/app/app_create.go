package app

import (
  "context"
  "fmt"
  "github.com/remicaumette/zaap.sh/pkg/api"
  "github.com/remicaumette/zaap.sh/pkg/protocol"
  "github.com/spf13/cobra"
)

func NewCreateCmd(client *api.Client) *cobra.Command {
  req := &protocol.CreateAppRequest{}
  cmd := &cobra.Command{
    Use:   "create [name]",
    Short: "Create an app",
    Long: `
This is the help create
`,
    Args: cobra.MinimumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
      req.Name = args[0]
      app, err := client.AppService.CreateApp(context.Background(), req)
      if err != nil {
        return err
      }
      fmt.Printf("%v\n", app.Name)
      return nil
    },
  }
  flags := cmd.Flags()
  flags.StringSliceVar(&req.Env, "env", []string{}, "Set environment variables")
  flags.StringVar(&req.Image, "image", "", "Set image")
  flags.Int32Var(&req.Memory, "memory", -1, "Memory limit")
  flags.Int32Var(&req.Cpu, "cpu", -1, "Number of CPUs")
  return cmd
}
