import client from './index'

export function fetchRunners() {
  return client.get('/runners')
    .then((res) => res.data.runners)
}

export function createRunner(payload) {
  return client.post('/runners', payload)
    .then((res) => res.data.runner)
}
