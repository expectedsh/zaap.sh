import React, { useState } from 'react'
import Head from 'next/head'
import Link from 'next/link'
import { GlobalLayout } from '~/layouts'
import { TextInput, Button } from '~/components'
import classnames from 'classnames/bind'
import styles from '~/assets/styles/login.scss'
import logo from '~/assets/images/logo.svg'
import githubLogo from '~/assets/images/githubLogo.svg'
import googleLogo from '~/assets/images/googleLogo.svg'
import { createContext } from 'vm'

const Login = () => {
  const [email, setEmail] = useState('')
  const cx = classnames.bind(styles)

  return (
    <GlobalLayout>
      <Head>
        <title>Login | Zaap</title>
      </Head>
      <div className={cx('root')}>
        <img src={logo} alt="Logo" title="Logo" className={cx('logo')} />
        <div className={cx('container')}>
          <div className={cx('title')}>Sign In</div>
          <form>
            <TextInput type="email" placeholder="Email Address" onChange={e => setEmail(e.target.value)}/>
            <Button type="submit" variant="success">
              Continue
            </Button>
          </form>
          <div className={cx('oauth-text')}>Or sign in with</div>
          <div className={cx('oauth-button-group')}>
            <div className={cx('oauth-button')}>
              <img src={githubLogo} title="Github Logo" alt="Github Logo" />
              Github
            </div>
            <div className={cx('oauth-button')}>
              <img src={googleLogo} title="Google Logo" alt="Google Logo" />
              Google
            </div>
          </div>
        </div>
        <div className={cx('alternative-link')}>
          Don’t have an account? <Link href="/sign_up">Sign Up</Link>.
        </div>
      </div>
    </GlobalLayout>
  )
}

export default Login
