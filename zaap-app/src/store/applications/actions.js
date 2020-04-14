import api from "~/utils/api"
import {
  FETCH_APPLICATIONS_PENDING,
  FETCH_APPLICATIONS_SUCCESS,
  FETCH_APPLICATIONS_ERROR,
  ADD_APPLICATION,
  UPDATE_APPLICATION,
  DELETE_APPLICATION,
  CREATE_APPLICATION_PENDING,
} from "./constants"
import { DELETE_APPLICATION_PENDING } from "~/store/application/constants"

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
    dispatch(createApplicationPending(true))
    return api.post("/applications", { name, image })
      .then(res => {
        dispatch(addApplication(res.data.application))
        return res.data.application
      })
      .finally(() => dispatch(createApplicationPending(false)))
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

export function updateApplication(payload) {
  return {
    type: UPDATE_APPLICATION,
    payload,
  }
}

export function deleteApplication(payload) {
  return {
    type: DELETE_APPLICATION,
    payload,
  }
}

export function createApplicationPending(payload) {
  return {
    type: CREATE_APPLICATION_PENDING,
    payload,
  }
}
