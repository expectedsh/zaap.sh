import React, { ReactNode, useState, useEffect } from "react"
import { Global, css } from "@emotion/core"
import { useLocation, useHistory } from "react-router-dom"
import { theme } from "utils/theme"
import { API_ENDPOINT } from "utils/api"
import { useThunkDispatch } from "hooks"
import { Box } from "components"
import * as Actions from "store/user/actions"
import * as S from "./styles"
import logo from "assets/images/logo.svg"
import githubLogo from "assets/images/githubLogo.svg"
import googleLogo from "assets/images/googleLogo.svg"

interface Props {
  title: string
  children: ReactNode
  alternative: ReactNode
  error?: Error
  loading?: boolean
}

const Layout = (props: Props) => {
  const history = useHistory()
  const dispatch = useThunkDispatch()
  const { search } = useLocation()
  const [loading, setLoading] = useState(props.loading)
  const [error, setError] = useState(props.error)

  const generateOAuthLink = (provider: string) =>
    `${API_ENDPOINT}/oauth/${provider}?redirect_url=${window.location.origin}/sign_in`

  useEffect(() => {
    let mounted = true
    const params = new URLSearchParams(search)
    const state = params.get("state")
    const code = params.get("code")
    if (!loading && state && code && (state === "github" || state === "google")) {
      setLoading(true)
      dispatch(Actions.authenticate({ provider: state, code }))
        .then(() => history.push("/"))
        .catch(error => {
          if (mounted) {
            setError(error)
          }
        })
        .finally(() => {
          if (mounted) {
            setLoading(false)
          }
        })
    }
    return () => {
      mounted = true
    }
  }, [dispatch, search, loading, history])

  return (
    <S.Root>
      <S.Logo src={logo} alt="Logo" title="Logo" />
      <S.Container>
        <S.Title>{props.title}</S.Title>
        {error && <Box variant="error" error={error} />}
        {props.children}
        <S.OAuthText>Or {props.title} with</S.OAuthText>
        <S.OAuthButtonGroup>
          <S.OAuthButton href={generateOAuthLink("github")}>
            <img src={githubLogo} title="Github Logo" alt="Github Logo" />
            Github
          </S.OAuthButton>
          <S.OAuthButton href={generateOAuthLink("google")}>
            <img src={googleLogo} title="Google Logo" alt="Google Logo" />
            Google
          </S.OAuthButton>
        </S.OAuthButtonGroup>
      </S.Container>
      <S.AlternativeLink>
        {props.alternative}
      </S.AlternativeLink>
      <Global styles={css`
        body {
          background: ${theme.color.blue};
        }
      `} />
    </S.Root>
  )
}

export default Layout
