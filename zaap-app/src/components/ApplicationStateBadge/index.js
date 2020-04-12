import React from 'react'
import PropTypes from 'prop-types'
import classnames from "classnames/bind"
import ApplicationStateName from "~/components/ApplicationStateName"
import style from "./ApplicationStateBadge.module.scss"

const cx = classnames.bind(style)

function getClassName(status) {
  switch (status) {
  case 'unknown':
    return "grey"
  case 'stopped':
    return "red"
  case 'starting':
  case 'running':
    return "green"
  }
}

function ApplicationStateBadge({ state }) {
  return (
    <div className={cx('root', getClassName(state))}>
      <ApplicationStateName state={state}/>
    </div>
  )
}

ApplicationStateBadge.propTypes = {
  state: PropTypes.string.isRequired,
}

export default ApplicationStateBadge
