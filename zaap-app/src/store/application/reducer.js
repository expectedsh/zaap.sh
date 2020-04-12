import {
  FETCH_APPLICATION_PENDING,
  FETCH_APPLICATION_SUCCESS,
  FETCH_APPLICATION_ERROR,
  DEPLOY_APPLICATION_PENDING,
  DEPLOY_APPLICATION_SUCCESS,
  DEPLOY_APPLICATION_ERROR,
} from './constants'

const initialState = {
  pending: false,
  application: null,
  error: null,
  deployPending: false,
  deployError: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case FETCH_APPLICATION_PENDING:
      return {
        ...state,
        pending: true,
      }
    case FETCH_APPLICATION_SUCCESS:
      return {
        ...state,
        pending: false,
        application: action.payload,
        error: null,
      }
    case FETCH_APPLICATION_ERROR:
      return {
        ...state,
        pending: false,
        application: null,
        error: action.error,
      }
    case DEPLOY_APPLICATION_PENDING:
      return {
        ...state,
        deployPending: true,
      }
    case DEPLOY_APPLICATION_SUCCESS:
      return {
        ...state,
        deployPending: false,
        application: action.payload,
        deployError: null,
      }
    case DEPLOY_APPLICATION_ERROR:
      return {
        ...state,
        deployPending: false,
        deployError: action.error,
      }
    default:
      return state
  }
}
