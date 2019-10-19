package server

import (
  "github.com/remicaumette/zaap.sh/pkg/models"
  "github.com/remicaumette/zaap.sh/pkg/util/github"
  "github.com/remicaumette/zaap.sh/pkg/util/httpx"
  "github.com/sirupsen/logrus"
  "golang.org/x/oauth2"
  "net/http"
)

func (s *Server) OAuthGithubRoute(ctx *httpx.Context) {
  ctx.Redirect(s.githubOAuthConfig.AuthCodeURL("", oauth2.AccessTypeOnline), http.StatusTemporaryRedirect)
}

func (s *Server) OAuthGithubCallbackRoute(ctx *httpx.Context) {
  token, err := s.githubOAuthConfig.Exchange(ctx.Context(), ctx.QueryParam("code"))
  if err != nil {
    ctx.ErrorBadRequest("Invalid oauth code.", nil)
    return
  }
  githubUser, err := github.GetUser(ctx.Context(), token)
  if err != nil {
    ctx.ErrorInternal(err)
    return
  }
  githubEmail, err := github.GetPrimaryEmail(ctx.Context(), token)
  if err != nil {
    ctx.ErrorInternal(err)
    return
  }
  user := &models.User{}
  if err := s.db.Where("github_id = ?", githubUser.ID).FirstOrCreate(user, models.User{
    Name:     githubUser.Name,
    Email:    githubEmail.Email,
    GithubID: githubUser.ID,
  }).Error; err != nil {
    ctx.ErrorInternal(err)
    return
  }
  logrus.Info(user)
}
