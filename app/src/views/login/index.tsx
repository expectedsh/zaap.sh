import React, { useState } from "react"
import { Global, css } from '@emotion/core'
import { Link } from "react-router-dom"
import { TextInput, Button } from 'components'
import { Root, Logo, Container, Title, OAuthText, OAuthButtonGroup, OAuthButton, AlternativeLink } from './styles'
import { theme } from "utils/theme"
import { API_ENDPOINT } from "utils/api"
import logo from 'assets/images/logo.svg'
import githubLogo from 'assets/images/githubLogo.svg'
import googleLogo from 'assets/images/googleLogo.svg'

const Login = () => {
  const [email, setEmail] = useState('')

  return (
    <>
      <Root>
        <Logo src={logo} alt="Logo" title="Logo" />
        <Container>
          <Title>Sign In</Title>
          <form>
            <TextInput type="email" placeholder="Email Address" onChange={e => setEmail(e.target.value)}/>
            <Button type="submit" variant="success">
              Continue
            </Button>
          </form>
          <OAuthText>Or sign in with</OAuthText>
          <OAuthButtonGroup>
            <OAuthButton href={`${API_ENDPOINT}/oauth/github`}>
              <img src={githubLogo} title="Github Logo" alt="Github Logo" />
              Github
            </OAuthButton>
            <OAuthButton href={`${API_ENDPOINT}/oauth/google`}>
              <img src={googleLogo} title="Google Logo" alt="Google Logo" />
              Google
            </OAuthButton>
          </OAuthButtonGroup>
        </Container>
        <AlternativeLink>
          Don’t have an account? <Link to="/sign_up">Sign Up</Link>.
        </AlternativeLink>
      </Root>
      <Global styles={css`
        body {
          background: ${theme.color.blue};
        }
      `}/>
    </>
  )
}

export default Login
