import React, { useEffect } from "react"
import { Link, useHistory } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import moment from "moment"
import { fetchApplications } from "~/store/applications/actions"
import Alert from "~/oldcomponents/Alert"
import Header from "~/oldcomponents/Header"
import ApplicationStatusBadge from "~/oldcomponents/ApplicationStatusBadge"
import Table from "~/oldcomponents/Table"
import style from "./ApplicationList.module.scss"
import { fetchRunners } from "~/store/runners/actions"

const cx = classnames.bind(style)

const tableConfig = runners => [
  {
    renderHeader: () => "Name",
    renderCell: app => app.name,
    cellClassName: cx("cell-name"),
  },
  {
    renderHeader: () => "Status",
    renderCell: app => <ApplicationStatusBadge status={app.status}/>,
    cellClassName: cx("cell-state"),
  },
  {
    renderHeader: () => "Runner",
    renderCell: app => {
      const runner = runners?.find(r => r.id === app.runnerId)
      return runner
        ? <Link to={`/runners`} className={cx('runner-link')}>{runner.name}</Link>
        : "Not found"
    },
    cellClassName: cx("cell-endpoint"),
  },
  {
    renderHeader: () => "Last update",
    renderCell: app => moment(app.updatedAt).fromNow(),
    cellClassName: cx("cell-created"),
  },
]

function ApplicationList() {
  const dispatch = useDispatch()
  const history = useHistory()
  const { pending, applications, error } = useSelector(state => state.applications)
  const runners = useSelector(state => state.runners.runners)

  useEffect(() => {
    dispatch(fetchApplications())
    dispatch(fetchRunners())
  }, [])

  function renderBody() {
    if (pending) {
      return <div>Loading...</div>
    }

    if (error) {
      return <Alert className="alert alert-error" error={error}/>
    }


    return applications ? (
      <Table
        config={tableConfig(runners)}
        dataSource={applications}
        onRowClick={app => history.push(`/apps/${app.id}`)}
        noData={
          <div className={cx("no-application")}>
            <div className={cx("title")}>
              You don't have application
            </div>
            <div className={cx("description")}>
              Create an application and it will show up here.
            </div>
          </div>
        }
      />
    ) : null
  }

  return (
    <>
      <Header preTitle="Overview" title="Applications">
        <Link className="btn btn-secondary" to="/apps/new">
          New application
        </Link>
      </Header>
      <div className="container">
        {renderBody()}
      </div>
    </>
  )
}

export default ApplicationList
