import api from '~/helpers/api'

export function list() {
  return api.get('/applications')
    .then((res) => res.data.applications)
}

export function create(payload) {
  return api.post('/applications', payload)
    .then((res) => res.data.application)
}
