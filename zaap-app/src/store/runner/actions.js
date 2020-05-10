import api from "~/utils/api"
import {
  FETCH_RUNNER_CLUSTER_ROLES_PENDING,
  FETCH_RUNNER_CLUSTER_ROLES_SUCCESS,
} from "~/store/runner/constants"

export function fetchClusterRoles({ id }) {
  return (dispatch) => {
    dispatch(fetchClusterRolesPending(true))
    return api.get(`/runners/${id}/cluster_roles`)
      .then(res => {
        const clusterRoles = res.data.clusterRoles
        dispatch(fetchClusterRolesSuccess(clusterRoles))
        return clusterRoles
      })
      .finally(() => dispatch(fetchClusterRolesPending(false)))
  }
}

export function fetchClusterRolesPending(payload) {
  return {
    type: FETCH_RUNNER_CLUSTER_ROLES_PENDING,
    payload,
  }
}

export function fetchClusterRolesSuccess(payload) {
  return {
    type: FETCH_RUNNER_CLUSTER_ROLES_SUCCESS,
    payload,
  }
}
