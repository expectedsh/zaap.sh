import React from "react"
import PropTypes from "prop-types"
import kubernetesLogo from '~/assets/images/kubernetes.png'
import dockerSwarmLogo from '~/assets/images/docker-swarm.png'

function getLogo(type) {
  switch (type) {
  case "kubernetes":
    return kubernetesLogo
  case "docker_swarm":
    return dockerSwarmLogo
  }
}

function RunnerTypeLogo({ type, ...props }) {
  const logo = getLogo(type)

  return logo ? <img src={logo} {...props} /> : null
}

RunnerTypeLogo.propTypes = {
  type: PropTypes.string.isRequired,
  className: PropTypes.string,
}

export default RunnerTypeLogo
