import React from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames/bind'
import style from './TextField.module.scss'

const cx = classnames.bind(style)

function TextField({ type = "text", name, label, value, ...props }) {
  return (
    <>
      {label && <label htmlFor={name}>{label}</label>}
      <input
        className={cx('root')} type={type} name={name} defaultValue={value} {...props} />
    </>
  )
}

TextField.propTypes = {
  type: PropTypes.string,
  name: PropTypes.string,
  placeholder: PropTypes.string,
  label: PropTypes.string,
  onChange: PropTypes.func,
  value: PropTypes.any,
}

export default TextField
