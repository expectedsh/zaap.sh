import React from 'react'
import { Link as RouterLink } from 'react-router-dom'

function Link({ children, ...props }) {
  return (
    <RouterLink {...props}>
      {children}
    </RouterLink>
  )
}

Link.propTypes = RouterLink.propTypes

export default Link
