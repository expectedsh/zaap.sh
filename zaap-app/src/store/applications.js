import { createSlice } from '@reduxjs/toolkit'

const applicationsSlice = createSlice({
  name: 'applications',
  initialState: {
    applications: null,
    currentApplication: null,
  },
  reducers: {
    setApplications(state, action) {
      state.applications = action.payload
    },
    setCurrentApplication(state, action) {
      state.currentApplication = action.payload
    },
  },
})

export const { setApplications, setCurrentApplication } = applicationsSlice.actions

export default applicationsSlice.reducer
