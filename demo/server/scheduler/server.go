package scheduler

import (
  "context"
  "errors"
  "github.com/remicaumette/zaap.sh/demo/protocol"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc/metadata"
)

type Scheduler struct {
  log *logrus.Entry
  tokens []string
  connections map[string]protocol.Scheduler_DeploymentEventsServer
}

func New() *Scheduler {
  return &Scheduler{
    log: logrus.WithField("app", "scheduler"),
    connections: make(map[string]protocol.Scheduler_DeploymentEventsServer),
  }
}

func (s *Scheduler) GetToken(ctx context.Context, r *protocol.GetTokenRequest) (*protocol.GetTokenResponse, error) {
  //token := uuid.NewV4().String()
  token := "2e73cbcc-b8ef-48e5-b1f5-dd0d75095c66"
  s.tokens = append(s.tokens, token)
  s.log.WithField("token", token).Info("new daemon registered")
  return &protocol.GetTokenResponse{
    Token: token,
  }, nil
}

func (s *Scheduler) RequestDeployment(token string, appName string, repository string) error {
  conn := s.connections[token]
  if conn == nil {
    return errors.New("connection not found")
  }
  return conn.Send(&protocol.DeploymentEventRequest{
    Type: protocol.DeploymentEventRequestType_DEPLOY_APP,
    RequestOneof: &protocol.DeploymentEventRequest_AppDeploymentRequest{
      AppDeploymentRequest:&protocol.AppDeploymentRequest{
        AppName: appName,
        Repository: repository,
      },
    },
  })
}

func mapType(t protocol.DeploymentEventResponseType) string {
  switch t {
  case protocol.DeploymentEventResponseType_OK:
    return "ok"
  case protocol.DeploymentEventResponseType_ERROR:
    return "error"
  }
  return ""
}

func (s *Scheduler) DeploymentEvents(stream protocol.Scheduler_DeploymentEventsServer) error {
  data, _ := metadata.FromIncomingContext(stream.Context())
  token := data.Get("token")[0]
  if token == "" {
    return errors.New("invalid token")
  }
  s.log.WithField("token", token).Info("daemon listening events")
  s.connections[token] = stream
  defer delete(s.connections, token)
  for {
    event, err := stream.Recv()
    if err != nil {
      s.log.WithError(err).Warn()
      break
    }
    logrus.
      WithField("token", token).
      WithField("type", mapType(event.Type)).
      WithField("message",  event.Message).
      Info("event from daemon received")
  }
  return nil
}
