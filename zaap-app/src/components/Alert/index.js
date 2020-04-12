import "./Alert.scss"
import React from "react"
import PropTypes from "prop-types"

function Alert({ error, children, ...props }) {
  return (
    <div {...props}>
      {error ? error.message : children}
    </div>
  )
}

Alert.propTypes = {
  error: PropTypes.object,
  className: PropTypes.string,
  onClick: PropTypes.func,
}

export default Alert
