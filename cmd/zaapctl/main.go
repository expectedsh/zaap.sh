package main

import (
  "github.com/remicaumette/zaap.sh/cmd/zaapctl/cmd"
  apiclient "github.com/remicaumette/zaap.sh/pkg/client"
  "github.com/sirupsen/logrus"
)

func main() {
  client, err := apiclient.New(":5200")
  if err != nil {
    logrus.Error(err)
  }
  defer client.Conn.Close()
  root := cmd.NewRootCmd(client)
  if err := root.Execute(); err != nil {
    logrus.Error(err)
  }
}
