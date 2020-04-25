import {
  FETCH_RUNNERS_PENDING,
  FETCH_RUNNERS_SUCCESS,
  FETCH_RUNNERS_ERROR,
  ADD_RUNNER,
  UPDATE_RUNNER,
  DELETE_RUNNER,
  CREATE_RUNNER_PENDING,
} from "./constants"

const initialState = {
  pending: false,
  runners: null,
  error: null,
  createPending: false,
}

export default function (state = initialState, action) {
  switch (action.type) {
  case FETCH_RUNNERS_PENDING:
    return {
      ...state,
      pending: true,
    }
  case FETCH_RUNNERS_SUCCESS:
    return {
      ...state,
      pending: false,
      runners: action.payload,
      error: null,
    }
  case FETCH_RUNNERS_ERROR:
    return {
      ...state,
      pending: false,
      error: action.error,
    }
  case ADD_RUNNER:
    return {
      ...state,
      runners: [
        ...(state.runners ?? []),
        action.payload,
      ],
    }
  case UPDATE_RUNNER:
    return {
      ...state,
      runners: [
        ...(state.runners ?? []).filter(v => v.id !== action.payload.id),
        action.payload,
      ],
    }
  case DELETE_RUNNER:
    return {
      ...state,
      runners: (state.runners ?? []).filter(v => v.id !== action.payload),
    }
  case CREATE_RUNNER_PENDING:
    return {
      ...state,
      createPending: action.payload,
    }
  default:
    return state
  }
}
