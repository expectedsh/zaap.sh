package main

import (
  "fmt"
  "github.com/remicaumette/zaap.sh/pkg/api"
  "github.com/remicaumette/zaap.sh/pkg/cmd"
  "github.com/sirupsen/logrus"
  "os"
)

func main() {
  addr := os.Getenv("ZAAP_ADDR")
  if addr == "" {
    addr = ":5200"
  }

  client, err := api.NewClient(addr)
  if err != nil {
    logrus.Error(err)
  }
  defer client.Conn.Close()

  root := cmd.NewRootCmd(client)
  if err := root.Execute(); err != nil {
    fmt.Fprint(os.Stderr, err)
    os.Exit(1)
  }
}
