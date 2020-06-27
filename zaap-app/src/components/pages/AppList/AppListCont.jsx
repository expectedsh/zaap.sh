import React, { useEffect, useState } from 'react'
import { fetchApplications } from '~/client/application'
import AppList from './AppList'

function AppListCont() {
  const [isLoading, setLoading] = useState(true)
  const [applications, setApplications] = useState([])
  const [error, setError] = useState(undefined)

  useEffect(() => {
    fetchApplications()
      .then((apps) => setApplications(apps))
      .catch((err) => setError(err))
      .finally(() => setLoading(false))
  }, [])

  return (
    <AppList
      loading={isLoading}
      applications={applications}
      error={error}
    />
  )
}

export default AppListCont
