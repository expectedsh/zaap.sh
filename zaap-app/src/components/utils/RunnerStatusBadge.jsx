import React from 'react'
import PropTypes from 'prop-types'
import styled from '@emotion/styled'
import { css } from '@emotion/core'
import RunnerStatusName from './RunnerStatusName'

function getColor(status) {
  switch (status) {
    case 'offline':
      return 'red'
    case 'online':
      return 'green'
    default:
      return 'grey'
  }
}

const StyledRunnerStatusBadge = styled.div((props) => css`
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

function RunnerStatusBadge({ status }) {
  return (
    <StyledRunnerStatusBadge color={getColor(status)}>
      <RunnerStatusName status={status} />
    </StyledRunnerStatusBadge>
  )
}

RunnerStatusBadge.propTypes = {
  status: PropTypes.string,
}

export default RunnerStatusBadge
