import React, { useState } from 'react'
import Head from 'next/head'
import Link from 'next/link'
import { GlobalLayout } from '~/layouts'
import { TextInput, Button } from '~/components'
import styles from '~/assets/styles/login.scss'
import logo from '~/assets/images/logo.svg'
import githubLogo from '~/assets/images/githubLogo.svg'
import googleLogo from '~/assets/images/googleLogo.svg'

const SignUp = () => {
  const [email, setEmail] = useState('')

  return (
    <GlobalLayout>
      <Head>
        <title>Sign Up | Zaap</title>
      </Head>
      <div className={styles.root}>
        <img src={logo} alt="Logo" title="Logo" className={styles.logo} />
        <div className={styles.container}>
          <div className={styles.title}>Sign Up</div>
          <form>
            <TextInput type="email" placeholder="Email Address" onChange={e => setEmail(e.target.value)} />
            <Button type="submit" variant="success">
              Continue
            </Button>
          </form>
          <div className={styles.oauthText}>Or sign up with</div>
          <div className={styles.oauthButtonGroup}>
            <div className={styles.oauthButton}>
              <img src={githubLogo} title="Github Logo" alt="Github Logo" />
              Github
            </div>
            <div className={styles.oauthButton}>
              <img src={googleLogo} title="Google Logo" alt="Google Logo" />
              Google
            </div>
          </div>
        </div>
        <div className={styles.needAccount}>
          Already have an account? <Link href="/login">Login</Link>.
        </div>
      </div>
    </GlobalLayout>
  )
}

export default SignUp
