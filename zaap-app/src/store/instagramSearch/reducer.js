import {INSTAGRAM_SEARCH_PENDING, INSTAGRAM_SEARCH_SUCCESS, INSTAGRAM_SEARCH_ERROR} from './constants'

const initialState = {
  pending: false,
  results: null,
  error: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case INSTAGRAM_SEARCH_PENDING:
      return {
        ...state,
        pending: true,
      }
    case INSTAGRAM_SEARCH_SUCCESS:
      return {
        ...state,
        pending: false,
        results: action.payload,
        error: null,
      }
    case INSTAGRAM_SEARCH_ERROR:
      return {
        ...state,
        pending: false,
        results: null,
        error: action.error,
      }
    default:
      return state
  }
}
