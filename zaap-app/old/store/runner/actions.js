import api from '~/utils/api';
import {
  FETCH_RUNNER_CLUSTER_ROLES_PENDING,
  FETCH_RUNNER_CLUSTER_ROLES_SUCCESS,
  FETCH_RUNNER_IMAGE_PULL_SECRETS_PENDING,
  FETCH_RUNNER_IMAGE_PULL_SECRETS_SUCCESS,
} from '~/store/runner/constants';

export function fetchClusterRoles({ id }) {
  return (dispatch) => {
    dispatch(fetchClusterRolesPending(true));
    return api.get(`/runners/${id}/cluster_roles`)
      .then((res) => {
        const { clusterRoles } = res.data;
        dispatch(fetchClusterRolesSuccess(clusterRoles));
        return clusterRoles;
      })
      .finally(() => dispatch(fetchClusterRolesPending(false)));
  };
}

export function fetchImagePullSecrets({ id }) {
  return (dispatch) => {
    dispatch(fetchImagePullSecretsPending(true));
    return api.get(`/runners/${id}/image_pull_secrets`)
      .then((res) => {
        const { imagePullSecrets } = res.data;
        dispatch(fetchImagePullSecretsSuccess(imagePullSecrets));
        return imagePullSecrets;
      })
      .finally(() => dispatch(fetchImagePullSecretsPending(false)));
  };
}

export function fetchClusterRolesPending(payload) {
  return {
    type: FETCH_RUNNER_CLUSTER_ROLES_PENDING,
    payload,
  };
}

export function fetchClusterRolesSuccess(payload) {
  return {
    type: FETCH_RUNNER_CLUSTER_ROLES_SUCCESS,
    payload,
  };
}

export function fetchImagePullSecretsPending(payload) {
  return {
    type: FETCH_RUNNER_IMAGE_PULL_SECRETS_PENDING,
    payload,
  };
}

export function fetchImagePullSecretsSuccess(payload) {
  return {
    type: FETCH_RUNNER_IMAGE_PULL_SECRETS_SUCCESS,
    payload,
  };
}
