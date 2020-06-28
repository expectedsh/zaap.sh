import { createSlice } from '@reduxjs/toolkit'

const runnersSlice = createSlice({
  name: 'runners',
  initialState: {
    runners: null,
    currentRunner: null,
  },
  reducers: {
    setRunners(state, action) {
      state.runners = action.payload
    },
    setCurrentRunner(state, action) {
      state.currentRunner = action.payload
    },
  },
})

export const { setRunners, setCurrentRunner } = runnersSlice.actions

export default runnersSlice.reducer
