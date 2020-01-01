package app

import (
  "context"
  "fmt"
  "github.com/remicaumette/zaap.sh/pkg/api"
  "github.com/remicaumette/zaap.sh/pkg/protocol"
  "github.com/spf13/cobra"
)

func NewCreateCmd(client *api.Client) *cobra.Command {
  opts := &protocol.CreateAppRequest{}
  cmd := &cobra.Command{
    Use:   "create",
    Short: "Create an app",
    Long: `
This is the help create
`,
    RunE: func(cmd *cobra.Command, args []string) error {
      app, err := client.AppService.CreateApp(context.Background(), opts)
      if err != nil {
        return err
      }
      fmt.Printf("%v\n", app.Name)
      return nil
    },
  }
  flags := cmd.Flags()
  flags.StringSliceVar(&opts.Env, "env", []string{}, "Set environment variables")
  flags.StringVar(&opts.Image, "image", "", "Set image")
  flags.Int32Var(&opts.Memory, "memory", -1, "Memory limit")
  flags.Int32Var(&opts.Cpu, "cpu", -1, "Number of CPUs")
  return cmd
}
