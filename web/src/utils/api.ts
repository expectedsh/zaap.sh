import axios from "axios"

export const API_ENDPOINT = process.env.ZAAP_API_URL || "http://localhost:3000"

export const client = axios.create({
  baseURL: API_ENDPOINT,
})
