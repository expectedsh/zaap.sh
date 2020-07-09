import PropTypes from 'prop-types'
import React from 'react'
import styled from '@emotion/styled'
import { css } from '@emotion/core'
import Heading from '~/components/atoms/Heading'
import Text from '~/components/atoms/Text'

const StyledAuthContainer = styled.div((props) => css`
  margin: 0 auto;
  padding: 32px;
  border-radius: 4px;
  max-width: 450px;
  background: ${props.theme.color.white};
`)

const StyledHeading = styled(Heading)`
  text-align: center;
  margin-bottom: 32px !important;
`

const StyledAlternative = styled(Text)((props) => css`
  margin: 24px 0;
  text-align: center;
  color: ${props.theme.color.white};

  a {
    color: ${props.theme.color.white} !important;
    text-decoration: underline;
  }
`)

function AuthContainer({ title, alternative, children }) {
  return (
    <>
      <StyledAuthContainer>
        {title && (
          <StyledHeading size="large">{title}</StyledHeading>
        )}
        {children}
      </StyledAuthContainer>
      {alternative && (
        <StyledAlternative>
          {alternative}
        </StyledAlternative>
      )}
    </>
  )
}

AuthContainer.propTypes = {
  title: PropTypes.string,
  alternative: PropTypes.node,
  children: PropTypes.node.isRequired,
}

export default AuthContainer
