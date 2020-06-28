import React, { lazy } from 'react'
import { Redirect, Route, Switch } from 'react-router'
import Dashboard from './Dashboard'

const AppListRoute = lazy(() => import('~/components/pages/AppList/AppListCont'))
const AppNewRoute = lazy(() => import('~/components/pages/AppNew/AppNewCont'))

const RunnerListRoute = lazy(() => import('~/components/pages/RunnerList/RunnerListCont'))

function DashboardCont() {
  return (
    <Dashboard>
      <Switch>
        <Route path="/apps/new" component={AppNewRoute} />
        <Route path="/apps" component={AppListRoute} />

        <Route path="/runners" component={RunnerListRoute} />

        <Redirect to="/apps" />
      </Switch>
    </Dashboard>
  )
}

export default DashboardCont
