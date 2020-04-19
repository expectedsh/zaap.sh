import React, { useEffect } from "react"
import { useDispatch, useSelector } from "react-redux"
import { fetchApplicationLogs } from "~/store/application/actions"

function Logs() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)

  useEffect(() => {
    dispatch(fetchApplicationLogs({ id: application.id }))
      .then((events) => {
        events.addEventListener("message", console.log)
      })
      .catch(console.error)
  }, [application])

  return (
    <div>Logs</div>
  )
}

export default Logs
