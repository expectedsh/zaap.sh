import React, { useEffect } from "react"
import { Redirect, Route, Switch, useParams } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { deployApplication, fetchApplication } from "~/store/application/actions"
import Alert from "~/oldcomponents/Alert"
import Button from "~/oldcomponents/Button"
import Header from "~/oldcomponents/Header"
import NavigationBar from "~/oldcomponents/NavigationBar"
import ApplicationOverview from "./ApplicationOverview"
import ApplicationSettings from "./ApplicationSettings"
import ApplicationLogs from "./ApplicationLogs"

function ApplicationShow() {
  const dispatch = useDispatch()
  const params = useParams()
  const { pending, application, error, deployPending } = useSelector(state => state.application)

  useEffect(() => {
    dispatch(fetchApplication({ id: params.id }))
  }, [params.id])

  function renderBody() {
    if (pending) {
      return <div>Loading...</div>
    }
    if (error) {
      return <Alert className="alert alert-error" error={error}/>
    }
    return application ? (
      <Switch>
        <Route path="/apps/:id/overview" component={ApplicationOverview}/>
        <Route path="/apps/:id/logs" component={ApplicationLogs}/>
        <Route path="/apps/:id/settings" component={ApplicationSettings}/>
        <Redirect from="/apps/:id" to={`/apps/${params.id}/overview`}/>
      </Switch>
    ) : null
  }

  function deploy() {
    if (application) {
      dispatch(deployApplication({ id: application.id }))
        .then(() => toast.success("Deployment requested"))
        .catch(err => toast.error(err.data?.message || err.response.statusText))
    }
  }

  return (
    <>
      <Header preTitle="Containers" title={application?.name ?? ""}>
        <Button className="btn btn-secondary" loading={deployPending} onClick={deploy}>
          Trigger deployment
        </Button>
      </Header>
      <NavigationBar style={{ marginTop: -32 }}>
        <NavigationBar.Link to={`/apps/${params.id}/overview`}>
          Overview
        </NavigationBar.Link>
        <NavigationBar.Link to={`/apps/${params.id}/logs`}>
          Logs
        </NavigationBar.Link>
        <NavigationBar.Link to={`/apps/${params.id}/settings`}>
          Settings
        </NavigationBar.Link>
      </NavigationBar>
      <div className="container">
        {renderBody()}
      </div>
    </>
  )
}

export default ApplicationShow
