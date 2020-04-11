import {INSTAGRAM_FETCH_POSTS_PENDING, INSTAGRAM_FETCH_POSTS_SUCCESS, INSTAGRAM_FETCH_POSTS_ERROR} from './constants'

const initialState = {
  pending: false,
  posts: null,
  pagination: null,
  error: null,
}

export default function (state = initialState, action) {
  switch (action.type) {
    case INSTAGRAM_FETCH_POSTS_PENDING:
      return {
        ...state,
        pending: true,
      }
    case INSTAGRAM_FETCH_POSTS_SUCCESS:
      const newPostIds = action.payload.posts.map(v => v.id)
      return {
        ...state,
        pending: false,
        posts: state.posts
          ? [
              ...state.posts.filter(v => !newPostIds.includes(v.id)),
              ...action.payload.posts
            ]
          : action.payload.posts,
        pagination: action.payload.pagination,
        error: null,
      }
    case INSTAGRAM_FETCH_POSTS_ERROR:
      return {
        ...state,
        pending: false,
        posts: null,
        pagination: null,
        error: action.error,
      }
    default:
      return state
  }
}
