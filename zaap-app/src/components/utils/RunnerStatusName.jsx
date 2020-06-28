import PropTypes from 'prop-types'

function RunnerStatusName({ status }) {
  switch (status) {
    case 'online':
      return 'Online'
    case 'offline':
      return 'Offline'
    default:
      return 'Unknown'
  }
}

RunnerStatusName.propTypes = {
  status: PropTypes.string,
}

export default RunnerStatusName
