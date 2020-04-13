import React, { useEffect } from 'react'
import { useParams } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { fetchApplication, deployApplication } from "~/store/application/actions"
import ApplicationStateBadge from "~/components/ApplicationStateBadge"
import Alert from "~/components/Alert"
import Button from "~/components/Button"
import { toast } from "react-toastify"

function ApplicationShow() {
  const dispatch = useDispatch()
  const params = useParams()
  const { pending, application, error, deployPending } = useSelector(state => state.application)

  useEffect(() => {
    dispatch(fetchApplication({ id: params.id }))
  }, [params])

  function renderBody() {
    if (pending) {
      return <div>Loading...</div>
    }
    if (error) {
      return <Alert className="alert alert-error" error={error} />
    }
    return application ? (
      <>
        <div>{application.name}</div>
        <ApplicationStateBadge state={application.state}/>
      </>
    ) : null
  }

  function deploy() {
    if (application) {
      dispatch(deployApplication({ id: application.id }))
        .then(() => toast.success('Deployment requested'))
        .catch(err => toast.error(err.response.statusText))
    }
  }

  return (
    <div className="container">
      <div className="header">
        <h1 className="header-title">
          Application
        </h1>
        <Button className="btn btn-success" loading={deployPending} onClick={deploy}>
          Request deployment
        </Button>
      </div>
      {renderBody()}
    </div>
  )
}

export default ApplicationShow
