import React from 'react'
import PropTypes from 'prop-types'

function ApplicationStateName({ state }) {
  switch (state) {
  case 'unknown':
    return "Unknown"
  case 'stopped':
    return "Stopped"
  case 'starting':
    return "Starting"
  case 'running':
    return "Running"
  default:
    return null
  }
}

ApplicationStateName.propTypes = {
  state: PropTypes.string.isRequired,
}

export default ApplicationStateName
