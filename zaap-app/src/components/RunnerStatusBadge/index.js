import React from "react"
import PropTypes from "prop-types"
import classnames from "classnames/bind"
import RunnerStatusName from "~/components/RunnerStatusName"
import style from "./RunnerStatusBadge.module.scss"

const cx = classnames.bind(style)

function getClassName(status) {
  switch (status) {
  case "offline":
    return "red"
  case "online":
    return "green"
  default:
    return "grey"
  }
}

function RunnerStatusBadge({ status }) {
  return (
    <div className={cx("root", getClassName(status))}>
      <RunnerStatusName state={status}/>
    </div>
  )
}

RunnerStatusBadge.propTypes = {
  status: PropTypes.string,
}

export default RunnerStatusBadge
