import React from 'react'
import { FORM_ERROR } from 'final-form'
import { useHistory } from 'react-router'
import { useDispatch } from 'react-redux'
import { register } from '~/client/auth'
import { setToken } from '~/store/authentication'
import AuthSignUp from './AuthSignUp'

function AuthSignUpCont() {
  const history = useHistory()
  const dispatch = useDispatch()

  function onSubmit(values) {
    return register(values)
      .then((token) => {
        dispatch(setToken(token))
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
