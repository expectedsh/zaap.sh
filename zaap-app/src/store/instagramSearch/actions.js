import api from '~/utils/api'
import { INSTAGRAM_SEARCH_PENDING, INSTAGRAM_SEARCH_SUCCESS, INSTAGRAM_SEARCH_ERROR } from './constants'

export function instagramSearchPending() {
  return {
    type: INSTAGRAM_SEARCH_PENDING
  }
}

export function instagramSearchSuccess(payload) {
  return {
    type: INSTAGRAM_SEARCH_SUCCESS,
    payload,
  }
}

export function instagramSearchError(error) {
  return {
    type: INSTAGRAM_SEARCH_ERROR,
    error,
  }
}

export function instagramSearch(query) {
  return dispatch => {
    dispatch(instagramSearchPending())
    return api.get(`/instagram/search?query=${encodeURIComponent(query)}`)
      .then(res => {
        dispatch(instagramSearchSuccess(res.data))
      })
      .catch(error => {
        dispatch(instagramSearchError(error))
      })
  }
}
