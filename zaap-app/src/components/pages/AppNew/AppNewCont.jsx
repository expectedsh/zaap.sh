import React, { useEffect } from 'react'
import { toast } from 'react-toastify'
import { useHistory } from 'react-router'
import { fetchRunners } from '~/client/runner'
import { createApplication } from '~/client/application'
import AppNew from './AppNew'

function AppNewCont() {
  const history = useHistory()

  useEffect(() => {
    fetchRunners()
      .catch(() => toast.error('Could not fetch runners.'))
  }, [])

  function onSubmit(values) {
    return createApplication({
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
