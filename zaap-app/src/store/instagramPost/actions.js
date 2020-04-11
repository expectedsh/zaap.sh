import api from '~/utils/api'
import {INSTAGRAM_FETCH_POST_PENDING, INSTAGRAM_FETCH_POST_SUCCESS, INSTAGRAM_FETCH_POST_ERROR} from './constants'

export function instagramFetchPost(id) {
  return dispatch => {
    dispatch(instagramFetchPostPending())
    return api.get(`/instagram/post/${id}`)
      .then(res => {
        dispatch(instagramFetchPostSuccess(res.data))
      })
      .catch(error => {
        dispatch(instagramFetchPostError(error))
      })
  }
}

export function instagramFetchPostPending() {
  return {
    type: INSTAGRAM_FETCH_POST_PENDING
  }
}

export function instagramFetchPostSuccess(payload) {
  return {
    type: INSTAGRAM_FETCH_POST_SUCCESS,
    payload,
  }
}

export function instagramFetchPostError(error) {
  return {
    type: INSTAGRAM_FETCH_POST_ERROR,
    error,
  }
}
