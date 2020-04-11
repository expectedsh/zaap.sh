import api from '~/utils/api'
import {INSTAGRAM_FETCH_POSTS_PENDING, INSTAGRAM_FETCH_POSTS_SUCCESS, INSTAGRAM_FETCH_POSTS_ERROR} from './constants'

export function instagramFetchPosts(userId, after) {
  return dispatch => {
    dispatch(instagramFetchPostsPending())
    return api.get(`/instagram/user/${userId}/posts?after=${encodeURIComponent(after)}`)
      .then(res => {
        dispatch(instagramFetchPostsSuccess(res.data))
      })
      .catch(error => {
        dispatch(instagramFetchPostsError(error))
      })
  }
}

export function instagramFetchMorePosts() {
  return (dispatch, getState) =>
    dispatch(instagramFetchPosts(
      getState().instagramUser.user.id,
      getState().instagramPosts.pagination.endCursor,
    ))
}

export function instagramFetchPostsPending() {
  return {
    type: INSTAGRAM_FETCH_POSTS_PENDING
  }
}

export function instagramFetchPostsSuccess(payload) {
  return {
    type: INSTAGRAM_FETCH_POSTS_SUCCESS,
    payload,
  }
}

export function instagramFetchPostsError(error) {
  return {
    type: INSTAGRAM_FETCH_POSTS_ERROR,
    error,
  }
}
