import React from 'react'
import classnames from 'classnames/bind'
import styles from './button.scss'

const Button = ({ variant, children, ...props }) => {
  const cx = classnames.bind(styles)
  return (
    <button {...props} className={cx('root', { success: variant === 'success' })}>
      {children}
    </button>
  )
}

export default Button
