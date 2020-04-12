import api from "~/utils/api"
import {
  FETCH_APPLICATIONS_PENDING,
  FETCH_APPLICATIONS_SUCCESS,
  FETCH_APPLICATIONS_ERROR,
  ADD_APPLICATION,
} from "./constants"

export function fetchApplications() {
  return dispatch => {
    dispatch(fetchApplicationsPending())
    return api.get("/applications")
      .then(res => {
        dispatch(fetchApplicationsSuccess(res.data.applications))
        return res.data.applications
      })
      .catch(error => {
        dispatch(fetchApplicationsError(error))
        return Promise.reject(error)
      })
  }
}

export function createApplication({ name, image }) {
  return dispatch => {
    return api.post("/applications", { name, image })
      .then(res => {
        dispatch(addApplication(res.data.application))
        return res.data.application
      })
  }
}

export function fetchApplicationsPending() {
  return {
    type: FETCH_APPLICATIONS_PENDING,
  }
}

export function fetchApplicationsSuccess(payload) {
  return {
    type: FETCH_APPLICATIONS_SUCCESS,
    payload,
  }
}

export function fetchApplicationsError(error) {
  return {
    type: FETCH_APPLICATIONS_ERROR,
    error,
  }
}

export function addApplication(payload) {
  return {
    type: ADD_APPLICATION,
    payload,
  }
}
