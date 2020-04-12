import api from "~/utils/api"
import { FETCH_USER_PENDING, FETCH_USER_SUCCESS, FETCH_USER_ERROR } from "./constants"

export function fetchUser() {
  return dispatch => {
    dispatch(fetchUserPending())
    return api.get("/me")
      .then(res => {
        dispatch(fetchUserSuccess(res.data.user))
        return res.data.user
      })
      .catch(error => {
        dispatch(fetchUserError(error))
        return Promise.reject(error)
      })
  }
}

export function fetchUserPending() {
  return {
    type: FETCH_USER_PENDING,
  }
}

export function fetchUserSuccess(payload) {
  return {
    type: FETCH_USER_SUCCESS,
    payload,
  }
}

export function fetchUserError(error) {
  return {
    type: FETCH_USER_ERROR,
    error,
  }
}
