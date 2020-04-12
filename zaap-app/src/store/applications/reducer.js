import {
  FETCH_APPLICATIONS_PENDING,
  FETCH_APPLICATIONS_SUCCESS,
  FETCH_APPLICATIONS_ERROR
} from './constants'

const initialState = {
  pending: false,
  applications: null,
  error: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case FETCH_APPLICATIONS_PENDING:
      return {
        ...state,
        pending: true,
      }
    case FETCH_APPLICATIONS_SUCCESS:
      return {
        ...state,
        pending: false,
        applications: action.payload,
        error: null,
      }
    case FETCH_APPLICATIONS_ERROR:
      return {
        ...state,
        pending: false,
        applications: null,
        error: action.error,
      }
    default:
      return state
  }
}
