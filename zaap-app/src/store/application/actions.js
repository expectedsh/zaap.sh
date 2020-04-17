import api, { ENDPOINT } from "~/utils/api"
import {
  updateApplication as updateApplicationFromList,
  deleteApplication as deleteApplicationFromList
} from "~/store/applications/actions"
import {
  FETCH_APPLICATION_PENDING,
  FETCH_APPLICATION_SUCCESS,
  FETCH_APPLICATION_ERROR,
  DEPLOY_APPLICATION_PENDING,
  DELETE_APPLICATION_PENDING,
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

export function deleteApplication({ id }) {
  return async dispatch => {
    dispatch(deleteApplicationPending(true))
    return api.delete(`/applications/${id}`)
      .then(() => {
        dispatch(deleteApplicationFromList(id))
      })
      .finally(() => dispatch(deleteApplicationPending(false)))
  }
}

export function deployApplication({ id }) {
  return dispatch => {
    dispatch(deployApplicationPending(true))
    return api.post(`/applications/${id}/deploy`)
      .then(res => {
        dispatch(updateApplicationFromList(res.data.application))
        return res.data.application
      })
      .finally(() => dispatch(deployApplicationPending(false)))
  }
}

export function fetchApplicationLogs({ id }) {
  return async (dispatch, getState) => {
    const token = getState().authentication.token
    return new EventSource(`${ENDPOINT}/applications/${id}/logs?token=${token}`)
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

export function deleteApplicationPending(payload) {
  return {
    type: DELETE_APPLICATION_PENDING,
    payload,
  }
}

export function deployApplicationPending(payload) {
  return {
    type: DEPLOY_APPLICATION_PENDING,
    payload,
  }
}
