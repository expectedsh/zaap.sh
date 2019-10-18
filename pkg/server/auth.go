package server

import (
  "github.com/remicaumette/zaap.sh/pkg/models"
  "github.com/remicaumette/zaap.sh/pkg/oauth/github"
  "github.com/remicaumette/zaap.sh/pkg/server/response"
  "github.com/sirupsen/logrus"
  "golang.org/x/oauth2"
  "net/http"
)

func (s *Server) OAuthGithubRoute(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, s.githubOAuthConfig.AuthCodeURL("", oauth2.AccessTypeOnline), http.StatusTemporaryRedirect)
}

func (s *Server) OAuthGithubCallbackRoute(w http.ResponseWriter, r *http.Request) {
  token, err := s.githubOAuthConfig.Exchange(r.Context(), r.FormValue("code"))
  if err != nil {
    response.ErrorBadRequest(w, "Invalid oauth code.", nil)
    return
  }
  githubUser, err := github.GetUser(r.Context(), token)
  if err != nil {
    logrus.WithError(err).Errorln("unable to retrieve your github data")
    response.ErrorInternal(w)
    return
  }
  githubEmail, err := github.GetPrimaryEmail(r.Context(), token)
  if err != nil {
    logrus.WithError(err).Errorln("unable to retrieve your github data")
    response.ErrorInternal(w)
    return
  }
  user := &models.User{}
  if err := s.db.Where("github_id = ?", githubUser.ID).FirstOrCreate(user, models.User{
    Name: githubUser.Name,
    Email: githubEmail.Email,
    GithubID: githubUser.ID,
  }).Error; err != nil {
    logrus.WithError(err).Errorln("unable to retrieve your github data")
    response.ErrorInternal(w)
    return
  }
  logrus.Info(user)
}
