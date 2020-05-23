import React, { useMemo } from "react"
import PropTypes from "prop-types"
import classnames from "classnames/bind"
import style from "./TextField.module.scss"

const cx = classnames.bind(style)

function TextField({ type = "text", input, meta, name, label, required, className, ...props }) {
  const fromArrayOrCurrent = s =>
    typeof s === "string" ? s : s?.[0]

  const error = useMemo(() => {
    return meta?.touched && (fromArrayOrCurrent(meta.error) || fromArrayOrCurrent(meta.submitError))
  }, [meta])

  return (
    <div className={cx("root", className)}>
      {label && (
        <label htmlFor={name}>
          {label}
          {required && <span className={cx("required")}>*</span>}
        </label>
      )}
      <input
        type={type}
        name={name}
        required={required}
        {...props}
        {...input}
      />
      {error && (
        <div className={cx("error")}>
          {error}
        </div>
      )}
    </div>
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
  required: PropTypes.bool,
  value: PropTypes.any,
  className: PropTypes.string,
}

export default TextField
