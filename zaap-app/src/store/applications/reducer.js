import {
  FETCH_APPLICATIONS_PENDING,
  FETCH_APPLICATIONS_SUCCESS,
  FETCH_APPLICATIONS_ERROR,
  ADD_APPLICATION,
  UPDATE_APPLICATION,
  DELETE_APPLICATION,
  CREATE_APPLICATION_PENDING,
} from "./constants"

const initialState = {
  pending: false,
  applications: null,
  error: null,
  createPending: false,
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
  case ADD_APPLICATION:
    return {
      ...state,
      applications: [
        ...(state.applications ?? []),
        action.payload,
      ],
    }
  case UPDATE_APPLICATION:
    return {
      ...state,
      applications: [
        ...(state.applications ?? []).filter(v => v.id !== action.payload.id),
        action.payload,
      ],
    }
  case DELETE_APPLICATION:
    return {
      ...state,
      applications: (state.applications ?? []).filter(v => v.id !== action.payload),
    }
  case CREATE_APPLICATION_PENDING:
    return {
      ...state,
      createPending: action.payload,
    }
  default:
    return state
  }
}
