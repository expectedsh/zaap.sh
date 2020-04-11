import api from '~/utils/api'
import {INSTAGRAM_FETCH_USER_PENDING, INSTAGRAM_FETCH_USER_SUCCESS, INSTAGRAM_FETCH_USER_ERROR} from './constants'
import {instagramFetchPostsPending, instagramFetchPostsSuccess} from "~/store/instagramPosts/actions";

export function instagramFetchUser(username) {
  return dispatch => {
    dispatch(instagramFetchUserPending())
    dispatch(instagramFetchPostsPending())
    return api.get(`/instagram/user/${username}`)
      .then(res => {
        dispatch(instagramFetchPostsSuccess({
          posts: res.data.posts,
          pagination: res.data.postPagination,
        }))
        dispatch(instagramFetchUserSuccess({
          ...res.data,
          posts: undefined,
          postPagination: undefined,
        }))
      })
      .catch(error => {
        dispatch(instagramFetchUserError(error))
      })
  }
}

export function instagramFetchUserPending() {
  return {
    type: INSTAGRAM_FETCH_USER_PENDING
  }
}

export function instagramFetchUserSuccess(payload) {
  return {
    type: INSTAGRAM_FETCH_USER_SUCCESS,
    payload,
  }
}

export function instagramFetchUserError(error) {
  return {
    type: INSTAGRAM_FETCH_USER_ERROR,
    error,
  }
}
