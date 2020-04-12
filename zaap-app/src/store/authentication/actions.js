import api from "~/utils/api"
import { AUTHENTICATION_PENDING, AUTHENTICATION_SUCCESS, AUTHENTICATION_ERROR } from "./constants"

export function register({ firstName, email, password }) {
  return dispatch => {
    dispatch(authenticationPending())
    return api.post('/users', { firstName, email, password })
      .then(res => {
        dispatch(authenticationSuccess(res.data.token))
        return res.data.token
      })
      .catch(error => {
        dispatch(authenticationError(error))
        return Promise.reject(error)
      })
  }
}

export function login({ email, password }) {
  return dispatch => {
    dispatch(authenticationPending())
    return api.post('/auth/login', { email, password })
      .then(res => {
        dispatch(authenticationSuccess(res.data.token))
        return res.data.token
      })
      .catch(error => {
        dispatch(authenticationError(error))
        return Promise.reject(error)
      })
  }
}

export function logout() {
  return authenticationSuccess(null)
}

export function authenticationPending() {
  return {
    type: AUTHENTICATION_PENDING,
  }
}

export function authenticationSuccess(payload) {
  return {
    type: AUTHENTICATION_SUCCESS,
    payload,
  }
}

export function authenticationError(error) {
  return {
    type: AUTHENTICATION_ERROR,
    error,
  }
}
