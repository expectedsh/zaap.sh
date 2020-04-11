import React from 'react'
import classnames from 'classnames/bind'
import WithBodyClasses from "~/components/WithBodyClasses"
import style from './Login.module.scss'

const cx = classnames.bind(style)

function Login() {
  return (
    <WithBodyClasses classNames={['login-page']}>
      <div className={cx('container')}>
        Login
      </div>
    </WithBodyClasses>
  )
}

export default Login
