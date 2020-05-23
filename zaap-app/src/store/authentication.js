import { createSlice } from '@reduxjs/toolkit'

const authenticationSlice = createSlice({
  name: 'authentication',
  initialState: {
    token: localStorage.getItem('token') ?? null,
  },
  reducers: {
    setToken(state, action) {
      state.token = action.payload.token
    },
  },
})

export const { setToken } = authenticationSlice.actions

export default authenticationSlice.reducer
