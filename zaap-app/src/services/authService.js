import api from '~/helpers/api'

export function login(payload) {
  return api.post('/auth/login', payload)
    .then((res) => res.data.token)
}
