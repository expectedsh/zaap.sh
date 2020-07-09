import api from '~/helpers/api'

export function findMe() {
  return api.get('/me')
    .then((res) => res.data.user)
}

export function updateMe(payload) {
  return api.patch('/me', payload)
    .then((res) => res.data.user)
}

export function create(payload) {
  return api.post('/users', payload)
    .then((res) => res.data.token)
}
