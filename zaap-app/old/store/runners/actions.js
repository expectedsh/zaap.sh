import api from '~/utils/api';
import {
  FETCH_RUNNERS_PENDING,
  FETCH_RUNNERS_SUCCESS,
  FETCH_RUNNERS_ERROR,
  ADD_RUNNER,
  UPDATE_RUNNER,
  DELETE_RUNNER,
  CREATE_RUNNER_PENDING,
} from './constants';

export function fetchRunners() {
  return (dispatch) => {
    dispatch(fetchRunnersPending());
    return api.get('/runners')
      .then((res) => {
        dispatch(fetchRunnersSuccess(res.data.runners));
        return res.data.runners;
      })
      .catch((error) => {
        dispatch(fetchRunnersError(error));
        return Promise.reject(error);
      });
  };
}

export function createRunner(payload) {
  return (dispatch) => {
    dispatch(createRunnerPending(true));
    return api.post('/runners', payload)
      .then((res) => {
        dispatch(addRunner(res.data.runner));
        return res.data.runner;
      })
      .finally(() => dispatch(createRunnerPending(false)));
  };
}

export function fetchRunnersPending() {
  return {
    type: FETCH_RUNNERS_PENDING,
  };
}

export function fetchRunnersSuccess(payload) {
  return {
    type: FETCH_RUNNERS_SUCCESS,
    payload,
  };
}

export function fetchRunnersError(error) {
  return {
    type: FETCH_RUNNERS_ERROR,
    error,
  };
}

export function addRunner(payload) {
  return {
    type: ADD_RUNNER,
    payload,
  };
}

export function updateRunner(payload) {
  return {
    type: UPDATE_RUNNER,
    payload,
  };
}

export function deleteRunners(payload) {
  return {
    type: DELETE_RUNNER,
    payload,
  };
}

export function createRunnerPending(payload) {
  return {
    type: CREATE_RUNNER_PENDING,
    payload,
  };
}
