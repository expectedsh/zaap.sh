import React from 'react'
import PropTypes from 'prop-types'
import styled from '@emotion/styled'
import { css } from '@emotion/core'
import ApplicationStatusName from './ApplicationStatusName'

function getColor(status) {
  switch (status) {
    case 'deploying':
    case 'running':
      return 'green'
    case 'crashed':
    case 'failed':
      return 'red'
    default:
      return 'grey'
  }
}

const StyledApplicationStatusBadge = styled.div((props) => css`
  padding: 4px 8px;
  width: fit-content;
  border-radius: 4px;
  font-family: ${props.theme.fontFamily.default};
  font-weight: ${props.theme.fontWeight.semiBold};
  font-size: ${props.theme.typography.text.small};
  text-transform: lowercase;
  color: ${props.theme.color.white};
  background: ${props.theme.color[props.color]['300']};
`)

function ApplicationStatusBadge({ status }) {
  return (
    <StyledApplicationStatusBadge color={getColor(status)}>
      <ApplicationStatusName status={status} />
    </StyledApplicationStatusBadge>
  )
}

ApplicationStatusBadge.propTypes = {
  status: PropTypes.string,
}

export default ApplicationStatusBadge
