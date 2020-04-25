import React, { useMemo } from 'react'
import Select from 'react-select'
import PropTypes from 'prop-types'
import classnames from 'classnames/bind'
import style from './SelectField.module.scss'

const cx = classnames.bind(style)

function SelectField({ input, meta, name, label, value, required, ...props }) {
  const fromArrayOrCurrent = s =>
    typeof s === 'string' ? s : s?.[0]

  const error = useMemo(() => {
    return meta?.touched && (fromArrayOrCurrent(meta.error) || fromArrayOrCurrent(meta.submitError))
  }, [meta])

  return (
    <div className={cx('root')}>
      {label && (
        <label htmlFor={name}>
          {label}
          {required && <span className={cx('required')}>*</span>}
        </label>
      )}
      <Select
        classNamePrefix="react-select"
        name={name}
        defaultValue={value}
        required={required}
        {...props}
        {...input}
      />
      {error && (
        <div className={cx('error')}>
          {error}
        </div>
      )}
    </div>
  )
}

SelectField.propTypes = {
  input: PropTypes.object,
  meta: PropTypes.object,
  name: PropTypes.string,
  placeholder: PropTypes.string,
  label: PropTypes.string,
  onChange: PropTypes.func,
  disabled: PropTypes.bool,
  required: PropTypes.bool,
  value: PropTypes.any,
  options: PropTypes.any,
}

export default SelectField
