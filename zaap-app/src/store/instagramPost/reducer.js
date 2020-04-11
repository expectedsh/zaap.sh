import {INSTAGRAM_FETCH_POST_PENDING, INSTAGRAM_FETCH_POST_SUCCESS, INSTAGRAM_FETCH_POST_ERROR} from './constants'

const initialState = {
  pending: false,
  post: null,
  error: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case INSTAGRAM_FETCH_POST_PENDING:
      return {
        ...state,
        pending: true,
      }
    case INSTAGRAM_FETCH_POST_SUCCESS:
      return {
        ...state,
        pending: false,
        post: action.payload,
        error: null,
      }
    case INSTAGRAM_FETCH_POST_ERROR:
      return {
        ...state,
        pending: false,
        post: null,
        error: action.error,
      }
    default:
      return state
  }
}
