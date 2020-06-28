import React, { lazy } from 'react'
import { Redirect, Route, Switch } from 'react-router'
import Dashboard from './Dashboard'

const AppListRoute = lazy(() => import('~/components/pages/AppList/AppListCont'))
const AppNewRoute = lazy(() => import('~/components/pages/AppNew/AppNewCont'))

const RunnerListRoute = lazy(() => import('~/components/pages/RunnerList/RunnerListCont'))
const RunnerNewRoute = lazy(() => import('~/components/pages/RunnerNew/RunnerNewCont'))

const SettingsRoute = lazy(() => import('~/components/pages/Settings/SettingsCont'))

function DashboardCont() {
  return (
    <Dashboard>
      <Switch>
        <Route path="/apps/new" component={AppNewRoute} />
        <Route path="/apps" component={AppListRoute} />

        <Route path="/runners/new" component={RunnerNewRoute} />
        <Route path="/runners" component={RunnerListRoute} />

        <Route path="/settings" component={SettingsRoute} />

        <Redirect to="/apps" />
      </Switch>
    </Dashboard>
  )
}

export default DashboardCont
