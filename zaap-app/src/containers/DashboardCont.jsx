import React from 'react'
import { Redirect, Route, Switch } from 'react-router'
import Dashboard from '~/components/pages/Dashboard'
import AppListCont from '~/containers/app/AppListCont'

function DashboardCont() {
  return (
    <Dashboard>
      <Switch>
        <Route path="/apps" component={AppListCont} />
        <Redirect to="/apps" />
      </Switch>
    </Dashboard>
  )
}

export default DashboardCont
