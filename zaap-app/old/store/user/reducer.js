import { FETCH_USER_PENDING, FETCH_USER_SUCCESS, FETCH_USER_ERROR } from './constants';

const initialState = {
  pending: false,
  user: null,
  error: null,
};

export default function (state = initialState, action) {
  switch (action.type) {
    case FETCH_USER_PENDING:
      return {
        ...state,
        pending: true,
      };
    case FETCH_USER_SUCCESS:
      return {
        ...state,
        pending: false,
        user: action.payload,
        error: null,
      };
    case FETCH_USER_ERROR:
      return {
        ...state,
        pending: false,
        user: null,
        error: action.error,
      };
    default:
      return state;
  }
}
