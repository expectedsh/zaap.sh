import {
  FETCH_RUNNER_CLUSTER_ROLES_PENDING,
  FETCH_RUNNER_CLUSTER_ROLES_SUCCESS,
  FETCH_RUNNER_IMAGE_PULL_SECRETS_PENDING,
  FETCH_RUNNER_IMAGE_PULL_SECRETS_SUCCESS,
} from "./constants"

const initialState = {
  clusterRoles: undefined,
  clusterRolesPending: false,
  imagePullSecrets: undefined,
  imagePullSecretsPending: false,
}

export default function (state = initialState, action) {
  switch (action.type) {
  case FETCH_RUNNER_CLUSTER_ROLES_PENDING:
    return {
      ...state,
      clusterRolesPending: action.payload,
    }
  case FETCH_RUNNER_CLUSTER_ROLES_SUCCESS:
    return {
      ...state,
      clusterRoles: action.payload,
    }
  case FETCH_RUNNER_IMAGE_PULL_SECRETS_PENDING:
    return {
      ...state,
      imagePullSecretsPending: action.payload,
    }
  case FETCH_RUNNER_IMAGE_PULL_SECRETS_SUCCESS:
    return {
      ...state,
      imagePullSecrets: action.payload,
    }
  default:
    return state
  }
}
