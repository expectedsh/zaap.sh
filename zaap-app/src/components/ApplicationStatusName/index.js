import React from "react"
import PropTypes from "prop-types"

function ApplicationStatusName({ status }) {
  switch (status) {
  case "deploying":
    return "Deploying"
  case "running":
    return "Running"
  case "crashed":
    return "Crashed"
  default:
    return "Unknown"
  }
}

ApplicationStatusName.propTypes = {
  status: PropTypes.string,
}

export default ApplicationStatusName
