import { createSlice } from '@reduxjs/toolkit'

const authenticationSlice = createSlice({
  name: 'authentication',
  initialState: {
    token: localStorage.getItem('token') ?? null,
  },
  reducers: {
    setToken(state, action) {
      const token = action.payload
      if (token) {
        localStorage.setItem('token', token)
      } else {
        localStorage.removeItem('token')
      }
      state.token = token
    },
  },
})

export const { setToken } = authenticationSlice.actions

export default authenticationSlice.reducer
