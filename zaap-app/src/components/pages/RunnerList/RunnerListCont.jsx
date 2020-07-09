import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { setRunners } from '~/store/runners'
import { runnerService } from '~/services'
import RunnerList from './RunnerList'

function RunnerListCont() {
  const dispatch = useDispatch()
  const runners = useSelector((s) => s.runners.runners)
  const [isLoading, setLoading] = useState(true)
  const [error, setError] = useState(undefined)

  useEffect(() => {
    runnerService.list()
      .then((fetchedRunners) => dispatch(setRunners(fetchedRunners)))
      .catch((err) => setError(err))
      .finally(() => setLoading(false))
  }, [])

  return (
    <RunnerList
      loading={isLoading}
      runners={runners}
      error={error}
    />
  )
}

export default RunnerListCont
