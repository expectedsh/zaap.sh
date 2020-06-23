import React, { useEffect } from 'react'
import { toast } from 'react-toastify'
import { fetchRunners } from '~/client/runner'
import AppNew from './AppNew'

function AppNewCont() {
  useEffect(() => {
    fetchRunners().catch(() => toast.error('Could not fetch runners.'))
  }, [])

  function onSubmit(values) {
    console.log(values)
    // return dispatch(createApplication({
    //   ...values,
    //   runnerId: values.runnerId,
    // }))
    //   .then(() => {
    //     toast.success('Application created.')
    //     history.push('/apps')
    //   })
    //   .catch((error) => {
    //     if (error.response.status === 422) {
    //       return error.data
    //     }
    //     toast.error(error.response.statusText)
    //   })
  }

  return (
    <AppNew isRunnerLoading runners={[]} onSubmit={onSubmit} />
  )
}

export default AppNewCont
