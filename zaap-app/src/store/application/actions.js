import api from "~/utils/api"
import {
  FETCH_APPLICATION_PENDING,
  FETCH_APPLICATION_SUCCESS,
  FETCH_APPLICATION_ERROR,
  DEPLOY_APPLICATION_ERROR,
  DEPLOY_APPLICATION_SUCCESS,
  DEPLOY_APPLICATION_PENDING,
} from "./constants"

export function fetchApplication({ id }) {
  return dispatch => {
    dispatch(fetchApplicationPending())
    return api.get(`/applications/${id}`)
      .then(res => {
        dispatch(fetchApplicationSuccess(res.data.application))
        return res.data.application
      })
      .catch(error => {
        dispatch(fetchApplicationError(error))
        return Promise.reject(error)
      })
  }
}

export function deployApplication({ id }) {
  return dispatch => {
    dispatch(deployApplicationPending())
    return api.post(`/applications/${id}/deploy`)
      .then(res => {
        dispatch(deployApplicationSuccess(res.data.application))
        return res.data.application
      })
      .catch(error => {
        dispatch(deployApplicationError(error))
        return Promise.reject(error)
      })
  }
}

export function fetchApplicationPending() {
  return {
    type: FETCH_APPLICATION_PENDING,
  }
}

export function fetchApplicationSuccess(payload) {
  return {
    type: FETCH_APPLICATION_SUCCESS,
    payload,
  }
}

export function fetchApplicationError(error) {
  return {
    type: FETCH_APPLICATION_ERROR,
    error,
  }
}

export function deployApplicationPending() {
  return {
    type: DEPLOY_APPLICATION_PENDING,
  }
}

export function deployApplicationSuccess(payload) {
  return {
    type: DEPLOY_APPLICATION_SUCCESS,
    payload,
  }
}

export function deployApplicationError(error) {
  return {
    type: DEPLOY_APPLICATION_ERROR,
    error,
  }
}
