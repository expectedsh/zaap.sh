import React, { useEffect } from 'react'
import PropTypes from 'prop-types'

function WithBodyClasses({ classNames, children = null }) {
  useEffect(() => {
    document.body.className = [
      ...document.body.className.split(" "),
      ...classNames
    ].join(" ")
    return () => {
      document.body.className = document.body.className
        .split(" ")
        .filter(className => classNames.includes(className))
        .join(" ")
    }
  }, [classNames])

  return children
}

WithBodyClasses.propTypes = {
  classNames: PropTypes.arrayOf(PropTypes.string),
  children: PropTypes.node,
}

export default WithBodyClasses
