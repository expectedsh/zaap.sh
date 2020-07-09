import axios from 'axios'
import store from '~/store'
import { camelizeKeys, decamelizeKeys } from './formatter'

export const ENDPOINT = process.env.API_ENDPOINT || 'http://localhost:3000'

const FORMATTER_OPTIONS = { exclude: (path) => path.includes('environment') }

const client = axios.create({
  baseURL: ENDPOINT,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
  transformRequest: [
    (data) => decamelizeKeys(data, FORMATTER_OPTIONS),
    ...axios.defaults.transformRequest,
  ],
})

client.interceptors.request.use((res) => {
  const { token } = store.getState().authentication
  if (token) {
    res.headers.Authorization = `Bearer ${token}`
  }
  return res
})

client.interceptors.response.use(
  (res) => {
    res.data = camelizeKeys(res.data, FORMATTER_OPTIONS)
    return res
  },
  (err) => {
    if (err.response.data) {
      err.data = camelizeKeys(err.response.data, FORMATTER_OPTIONS)
    }
    return Promise.reject(err)
  },
)

export default client
