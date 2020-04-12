import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames/bind'
import style from './TextField.scss'

const cx = classnames.bind(style)

function TextField({ type = "text", input, meta, name, label, value, ...props }) {
  const error = useMemo(() => {
    return meta.touched && (meta.error || meta.submitError)
  }, [meta])

  return (
    <>
      {label && <label htmlFor={name}>{label}</label>}
      <input
        className={cx('root')}
        type={type}
        name={name}
        defaultValue={value}
        {...props}
        {...input}
      />
      {error && (
        <div className={cx('error')}>
          {error}
        </div>
      )}
    </>
  )
}

TextField.propTypes = {
  type: PropTypes.string,
  input: PropTypes.object,
  meta: PropTypes.object,
  name: PropTypes.string,
  placeholder: PropTypes.string,
  label: PropTypes.string,
  onChange: PropTypes.func,
  disabled: PropTypes.bool,
  value: PropTypes.any,
}

export default TextField
