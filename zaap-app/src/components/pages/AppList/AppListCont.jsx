import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { setApplications } from '~/store/applications'
import { setRunners } from '~/store/runners'
import { applicationService, runnerService } from '~/services'
import AppList from './AppList'

function AppListCont() {
  const dispatch = useDispatch()
  const applications = useSelector((s) => s.applications.applications)
  const runners = useSelector((s) => s.runners.runners)
  const [isLoading, setLoading] = useState(true)
  const [error, setError] = useState(undefined)

  useEffect(() => {
    Promise.all([
      runnerService.list(),
      applicationService.list(),
    ])
      .then(([fetchedRunners, fetchedApps]) => {
        dispatch(setRunners(fetchedRunners))
        dispatch(setApplications(fetchedApps))
      })
      .catch((err) => setError(err))
      .finally(() => setLoading(false))
  }, [])

  return (
    <AppList
      loading={isLoading}
      applications={applications}
      runners={runners}
      error={error}
    />
  )
}

export default AppListCont
