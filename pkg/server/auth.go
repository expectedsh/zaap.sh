package server

import (
  "github.com/remicaumette/zaap.sh/pkg/models"
  "github.com/remicaumette/zaap.sh/pkg/util/github"
  "github.com/remicaumette/zaap.sh/pkg/util/httpx"
  "golang.org/x/oauth2"
  "net/http"
)

func (s *Server) OAuthGithubRoute(ctx *httpx.Context) {
  if redirect := ctx.QueryParam("redirect_url"); redirect != "" {
    s.GithubOAuthConfig.RedirectURL = redirect
  }
  ctx.Redirect(s.GithubOAuthConfig.AuthCodeURL("github", oauth2.AccessTypeOnline), http.StatusTemporaryRedirect)
}

type AuthenticateRequest struct {
  Provider string `json:"provider"`
  Code     string `json:"code"`
  Email    string `json:"email"`
}

func (s *Server) AuthenticateRoute(ctx *httpx.Context, req AuthenticateRequest) {
  if req.Provider == "github" {
    token, err := s.GithubOAuthConfig.Exchange(ctx.Context(), req.Code)
    if err != nil {
      ctx.ErrorBadRequest("Invalid code.", nil)
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
    if err := s.DB.Where("github_id = ?", githubUser.ID).FirstOrCreate(user, models.User{
      Name:     githubUser.Name,
      Email:    githubEmail.Email,
      GithubID: githubUser.ID,
    }).Error; err != nil {
      if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
        ctx.ErrorBadRequest("Existing account found for your email. Login with email and connect your Github account.", nil)
      } else {
        ctx.ErrorInternal(err)
      }
      return
    }
    jwt, err := user.NewToken()
    if err != nil {
      ctx.ErrorInternal(err)
      return
    }
    ctx.Json(http.StatusOK, map[string]interface{}{
      "token": jwt,
    })
  } else {
    ctx.ErrorBadRequest("Invalid provider.", nil)
  }
}
