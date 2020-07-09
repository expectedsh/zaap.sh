import React from 'react'
import { toast } from 'react-toastify'
import { useHistory } from 'react-router'
import { runnerService } from '~/services'
import RunnerNew from './RunnerNew'

function RunnerNewCont() {
  const history = useHistory()

  function onSubmit(values) {
    return runnerService.create(values)
      .then(() => {
        toast.success('Runner registered.')
        history.push('/runners')
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
    <RunnerNew onSubmit={onSubmit} />
  )
}

export default RunnerNewCont
