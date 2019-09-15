import React from 'react'
import { DashboardLayout } from '~/layouts'
import { PageHeader, RouteCard } from '~/components'

const Home = () => (
  <DashboardLayout>
    <div className="container">
      <PageHeader title="Overview" />
      <div className="row">
        <div className="col-md-6">
          <RouteCard />
        </div>
        <div className="col-md-6">
          <RouteCard online />
        </div>
        <div className="col-md-6">
          <RouteCard />
        </div>
      </div>
    </div>
  </DashboardLayout>
)

export default Home
