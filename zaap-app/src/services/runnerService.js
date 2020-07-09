import api from '~/helpers/api'

export function list() {
  return api.get('/runners')
    .then((res) => res.data.runners)
}

export function create(payload) {
  return api.post('/runners', payload)
    .then((res) => res.data.runner)
}
