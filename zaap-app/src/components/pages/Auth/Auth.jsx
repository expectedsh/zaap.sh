import React from 'react'
import PropTypes from 'prop-types'
import styled from '@emotion/styled'
import { Global, css } from '@emotion/core'
import theme from '~/style/theme'
import background from '~/assets/images/login-background.svg'
import logo from '~/assets/images/logo.svg'

const globalStyles = css`
  body {
    background: ${theme.color.primary['300']} url(${background}) center;
  }
`

const StyledRoot = styled.div`
  margin: 42px 0;
`

const StyledLogo = styled.img`
  display: block;
  margin: 0 auto 32px;
  height: 52px;
`

function Auth({ children }) {
  return (
    <StyledRoot>
      <Global styles={globalStyles} />
      <StyledLogo src={logo} alt="Zaap logo" />
      {children}
    </StyledRoot>
  )
}

Auth.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Auth
