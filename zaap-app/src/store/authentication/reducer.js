import { AUTHENTICATION_PENDING, AUTHENTICATION_SUCCESS, AUTHENTICATION_ERROR } from "./constants"

const initialState = {
  pending: false,
  token: localStorage.getItem('token') ?? null,
  error: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
  case AUTHENTICATION_PENDING:
    return {
      ...state,
      pending: true,
    }
  case AUTHENTICATION_SUCCESS:
    if (action.payload) {
      localStorage.setItem('token', action.payload)
    } else {
      localStorage.removeItem('token')
    }
    return {
      ...state,
      pending: false,
      token: action.payload,
      error: null,
    }
  case AUTHENTICATION_ERROR:
    return {
      ...state,
      pending: false,
      error: action.error,
    }
  default:
    return state
  }
}
