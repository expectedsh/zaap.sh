import PropTypes from 'prop-types'
import React from 'react'
import styled from '@emotion/styled'
import { css } from '@emotion/core'
import { Link as RouterLink } from 'react-router-dom'

const StyledLink = styled(RouterLink)((props) => {
  const typo = props.theme.typography.text[props.size]

  return css`
    font-family: ${props.theme.fontFamily.default};
    font-size: ${typo.fontSize};
    line-height: ${typo.lineHeight};
    
    :focus {
      outline: none;
    }
  `
})

function Link({ children, ...props }) {
  return (
    <StyledLink {...props}>
      {children}
    </StyledLink>
  )
}

Link.propTypes = {
  size: PropTypes.oneOf(['small', 'medium', 'large']),
  to: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
}

Link.defaultProps = {
  size: 'medium',
}

export default Link
