import React, { lazy } from 'react'
import { Redirect, Route, Switch } from 'react-router'
import Dashboard from './Dashboard'

const AppListRoute = lazy(() => import('~/components/pages/AppList/AppListCont'))
const AppNewRoute = lazy(() => import('~/components/pages/AppNew/AppNewCont'))

function DashboardCont() {
  return (
    <Dashboard>
      <Switch>
        <Route path="/apps/new" component={AppNewRoute} />
        <Route path="/apps" component={AppListRoute} />
        <Redirect to="/apps" />
      </Switch>
    </Dashboard>
  )
}

export default DashboardCont
