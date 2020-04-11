import axios from 'axios'

export default axios.create({
  baseURL: process.env.API_ENDPOINT || 'http://localhost:3000/',
})
