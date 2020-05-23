import React from 'react'
import { useSelector } from 'react-redux'
import DashboardCont from '~/containers/DashboardCont'
import AuthCont from '~/containers/auth/AuthCont'

function AppCont() {
  const token = useSelector((state) => state.authentication.token)

  return token ? <DashboardCont /> : <AuthCont />
}

export default AppCont
