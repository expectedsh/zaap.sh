import axios from "axios"
import humps from "humps"
import store from "~/store"

export const ENDPOINT = process.env.API_ENDPOINT || "http://localhost:3000"

const client = axios.create({
  baseURL: ENDPOINT,
  headers: {
    "Content-Type": "application/json",
    "Accept": "application/json",
  },
  transformRequest: [
    data => humps.decamelizeKeys(data),
    ...axios.defaults.transformRequest,
  ],
})

client.interceptors.request.use(res => {
  const token = store.getState().authentication.token
  if (token) {
    res.headers["Authorization"] = `Bearer ${token}`
  }
  return res
})

client.interceptors.response.use(
  (res) => {
    res.data = humps.camelizeKeys(res.data)
    return res
  },
  (err) => {
    if (err.response.data) {
      err.data = humps.camelizeKeys(err.response.data)
    }
    return Promise.reject(err)
  },
)

export default client
