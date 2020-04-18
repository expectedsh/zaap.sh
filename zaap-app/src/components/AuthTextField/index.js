import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames/bind'
import style from './AuthTextField.module.scss'

const cx = classnames.bind(style)

function AuthTextField({ type = "text", input, meta, name, value, ...props }) {
  const error = useMemo(() => {
    return meta?.touched && (meta.error || meta.submitError)
  }, [meta])

  return (
    <>
      <input
        className={cx('root', { 'has-error': error != null })}
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

AuthTextField.propTypes = {
  type: PropTypes.string,
  input: PropTypes.object,
  meta: PropTypes.object,
  name: PropTypes.string,
  placeholder: PropTypes.string,
  onChange: PropTypes.func,
  value: PropTypes.any,
}

export default AuthTextField
