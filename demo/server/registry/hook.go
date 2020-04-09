package registry

import (
  "encoding/json"
  scheduler2 "github.com/remicaumette/zaap.sh/demo/server/scheduler"
  "github.com/sirupsen/logrus"
  "io/ioutil"
  "net/http"
  "regexp"
)

var repositoryRegex = regexp.MustCompile(`^([a-zA-Z0-9\-]+)/([a-zA-Z0-9\-]+)$`)

type HookServer struct {
  log *logrus.Entry
  scheduler *scheduler2.Scheduler
}

func NewHookServer(scheduler *scheduler2.Scheduler) *HookServer {
  return &HookServer{
    log: logrus.WithField("app", "hook-server"),
    scheduler: scheduler,
  }
}

func (s *HookServer) Handler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    s.log.WithError(err).Fatal("failed to read the body")
    return
  }
  payload := make(map[string]interface{})
  if err = json.Unmarshal(body, &payload); err != nil {
    s.log.WithError(err).Fatal("could not parse the body")
    return
  }
  event := payload["events"].([]interface{})[0].(map[string]interface{})
  if event["action"] == "push" {
    s.log.Info("push action received")
    repository := event["target"].(map[string]interface{})["repository"].(string)
    found := repositoryRegex.FindStringSubmatch(repository)
    if len(found) != 3 {
      s.log.WithField("repository", repository).Warn("invalid repository name, could not deploy")
      w.WriteHeader(200)
      w.Write([]byte("ok"))
      return
    }
    schedulerToken := found[1]
    appName := found[2]
    log := s.log.
      WithField("repository", repository).
      WithField("scheduler", schedulerToken).
      WithField("app_name", appName)
    log.Info("dispatching deployment request to the scheduler")
    if err := s.scheduler.RequestDeployment(schedulerToken, appName, repository); err != nil {
      log.WithError(err).Error("failed to request deployment")
    }
  }
  w.WriteHeader(200)
  w.Write([]byte("ok"))
}

func (s *HookServer) Start() error {
  httpServer := http.Server{
    Addr:    ":8091",
    Handler: http.HandlerFunc(s.Handler),
  }
  s.log.Infof("listening on %v", httpServer.Addr)
  return httpServer.ListenAndServe()
}
