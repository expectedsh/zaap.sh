import React from 'react'
import { FORM_ERROR } from 'final-form'
import { useHistory } from 'react-router'
import { login } from '~/client/auth'
import { setToken } from '~/store/authentication'
import AuthSignIn from './AuthSignIn'

function AuthSignInCont() {
  const history = useHistory()

  function onSubmit(values) {
    return login(values)
      .then((token) => {
        setToken(token)
        history.push('/')
      })
      .catch((error) => {
        if (error.response.status === 422) {
          return error.data
        }
        if (error.response.status === 404) {
          return { [FORM_ERROR]: error.data.message }
        }
        return { [FORM_ERROR]: error.response.statusText }
      })
  }

  return (
    <AuthSignIn onSubmit={onSubmit} />
  )
}

export default AuthSignInCont
