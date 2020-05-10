import {
  FETCH_RUNNER_CLUSTER_ROLES_PENDING,
  FETCH_RUNNER_CLUSTER_ROLES_SUCCESS,
} from "./constants"

const initialState = {
  clusterRoles: undefined,
  clusterRolesPending: false,
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
  default:
    return state
  }
}
