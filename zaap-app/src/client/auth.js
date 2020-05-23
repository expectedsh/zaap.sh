import client from './index'

export function login({ email, password }) {
  return client.post('/auth/login', { email, password })
    .then((res) => res.data.token)
}

export function register({ firstName, email, password }) {
  return client.post('/users', { firstName, email, password })
    .then((res) => res.data.token)
}
