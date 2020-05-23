import React, { useMemo } from 'react'
import Select from 'react-select'
import PropTypes from 'prop-types'
import classnames from 'classnames/bind'
import style from './SelectField.module.scss'

const cx = classnames.bind(style)

function SelectField({ input, meta, name, label, required, ...props }) {
  const fromArrayOrCurrent = s =>
    typeof s === 'string' ? s : s?.[0]

  const value = useMemo(() => {
    const curr = input.value || value
    if (!curr || !props.options) {
      return undefined
    }
    return props.isMulti
      ? curr.map(v => props.options.find(o => o.value === v)).filter(v => !!v)
      : props.options.find(o => o.value === curr)
  }, [input?.value, props.value, props.options, props.isMulti])

  const error = useMemo(() => {
    return meta?.touched && (fromArrayOrCurrent(meta.error) || fromArrayOrCurrent(meta.submitError))
  }, [meta])

  function onChange(event) {
    const handler = input.onChange || props.onChange
    handler?.(props.isMulti ? (event?.map(v => v.value) ?? []) : event?.value)
  }

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
        required={required}
        {...props}
        {...input}
        onChange={onChange}
        value={value}
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
  isMulti: PropTypes.bool,
  loadOptions: PropTypes.func,
  value: PropTypes.any,
  options: PropTypes.any,
}

export default SelectField
