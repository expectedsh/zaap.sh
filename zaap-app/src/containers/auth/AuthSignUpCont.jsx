import React from 'react'
import { FORM_ERROR } from 'final-form'
import { useHistory } from 'react-router'
import { register } from '~/client/auth'
import { setToken } from '~/store/authentication'
import AuthSignUp from '~/components/pages/auth/AuthSignUp'

function AuthSignUpCont() {
  const history = useHistory()

  function onSubmit(values) {
    return register(values)
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
    <AuthSignUp onSubmit={onSubmit} />
  )
}

export default AuthSignUpCont
