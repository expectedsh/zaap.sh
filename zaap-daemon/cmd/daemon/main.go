package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/internal/daemon"
	"github.com/remicaumette/zaap.sh/pkg/configs"
)

// todo: check if the connection is always up, if not retry it until it is alive
//       don't forget to store locally all messages that are not sent

func main() {

	logrus.Info("Starting 'daemon' ...")

	daemonConfig := configs.Daemon{}
	if err := envconfig.Process("", &daemonConfig); err != nil {
		logrus.Panic(err)
	}

	daemonProxy := url.URL{Scheme: "ws", Host: daemonConfig.DaemonProxyAddress, Path: "/daemon"}

	connection, _, err := websocket.DefaultDialer.Dial(daemonProxy.String(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	connection.SetCloseHandler(func(code int, text string) error {
		fmt.Println(code, text)
		return nil
	})

	stop := make(chan os.Signal, 1)

	go func() {
		<-stop
		logrus.Info("Stopping the connection")
		connection.Close()
	}()

	go daemon.ProxySocketHandler(connection)

	if dockerListener, err := daemon.NewDockerEventListener(daemonConfig); err != nil {
		logrus.Panic(err)
	} else {
		dockerListener.Listen()
	}

}
