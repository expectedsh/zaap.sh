import React, { useEffect } from 'react'
import { toast } from 'react-toastify'
import { useHistory } from 'react-router'
import { applicationService, runnerService } from '~/services'
import AppNew from './AppNew'

function AppNewCont() {
  const history = useHistory()

  useEffect(() => {
    runnerService.list()
      .catch(() => toast.error('Could not fetch runners.'))
  }, [])

  function onSubmit(values) {
    return applicationService.create({
      ...values,
      runnerId: values.runnerId,
    })
      .then(() => {
        toast.success('Application created.')
        history.push('/apps')
      })
      .catch((error) => {
        if (error.response.status === 422) {
          return error.data
        }
        toast.error(error.response.statusText)
        return undefined
      })
  }

  return (
    <AppNew isRunnerLoading runners={[]} onSubmit={onSubmit} />
  )
}

export default AppNewCont
