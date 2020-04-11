import React from 'react'
import { useForm, Controller } from "react-hook-form"
import classnames from 'classnames/bind'
import WithBodyClasses from "~/components/WithBodyClasses"
import logo from '~/assets/images/logo.svg'
import style from './SignUp.module.scss'
import TextField from "~/components/TextField"

const cx = classnames.bind(style)

function SignUp() {
  const { errors, setError, handleSubmit, control } = useForm()

  function onSubmit(values) {
    console.log(values)
  }

  return (
    <WithBodyClasses classNames={['login-page']}>
      <div className={cx('root')}>
        <img className={cx('logo')} src={logo} alt="Zaap logo"/>
        <div className={cx('container')}>
          <h1 className={cx('title')}>Sign Up</h1>
          <form onSubmit={handleSubmit(onSubmit)}>
            <Controller
              name="firstName"
              control={control}
              as={<TextField name="firstName" placeholder="First name"/>}
            />
            <Controller
              name="email"
              control={control}
              as={<TextField type="email" name="email" placeholder="Email"/>}
            />
            <Controller
              name="password"
              control={control}
              as={<TextField type="password" name="password" placeholder="Password"/>}
            />
            <button className="btn btn-success" type="submit">
              Continue
            </button>
          </form>
        </div>
      </div>
    </WithBodyClasses>
  )
}

export default SignUp
