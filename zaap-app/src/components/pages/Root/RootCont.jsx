import React, { Suspense } from 'react'
import { useSelector } from 'react-redux'
import DashboardCont from '~/components/pages/Dashboard/DashboardCont'
import AuthCont from '~/components/pages/Auth/AuthCont'

function RootCont() {
  const token = useSelector((state) => state.authentication.token)

  return (
    <Suspense fallback={<p>error</p>}>
      {token ? <DashboardCont /> : <AuthCont />}
    </Suspense>
  )
}

export default RootCont
