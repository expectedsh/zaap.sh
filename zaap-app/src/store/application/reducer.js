import {
  FETCH_APPLICATION_PENDING,
  FETCH_APPLICATION_SUCCESS,
  FETCH_APPLICATION_ERROR,
  DEPLOY_APPLICATION_PENDING,
  UPDATE_APPLICATION_PENDING,
  DELETE_APPLICATION_PENDING,
} from './constants'

const initialState = {
  pending: false,
  application: null,
  error: null,
  deployPending: false,
  updatePending: false,
  deletePending: false,
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
        application: state.application?.id === action.payload.id ? {
          ...state.application,
          ...action.payload,
        } : action.payload,
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
        deployPending: action.payload,
      }
    case UPDATE_APPLICATION_PENDING:
      return {
        ...state,
        updatePending: action.payload,
      }
    case DELETE_APPLICATION_PENDING:
      return {
        ...state,
        deletePending: action.payload,
      }
    default:
      return state
  }
}
