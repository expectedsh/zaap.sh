import client from './index'

export function fetchApplications() {
  return client.get('/applications')
    .then((res) => res.data.applications)
}

export function createApplication(payload) {
  return client.post('/applications', payload)
    .then((res) => res.data.application)
}
