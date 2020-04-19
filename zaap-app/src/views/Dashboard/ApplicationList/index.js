import React, { useEffect } from "react"
import { Link, useHistory } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import moment from "moment"
import { fetchApplications } from "~/store/applications/actions"
import Alert from "~/components/Alert"
import Header from "~/components/Header"
import ApplicationStateBadge from "~/components/ApplicationStateBadge"
import Table from "~/components/Table"
import style from "./ApplicationList.module.scss"

const cx = classnames.bind(style)

const tableConfig = [
  {
    renderHeader: () => "Name",
    renderCell: () => "My name",
    cellClassNames: cx("cell-name"),
  },
  {
    renderHeader: () => "Status",
    renderCell: app => <ApplicationStateBadge state={app.state}/>,
    cellClassNames: cx("cell-state"),
  },
  {
    renderHeader: () => "Endpoint",
    renderCell: () => "My name",
    cellClassNames: cx("cell-endpoint"),
  },
  {
    renderHeader: () => "Created",
    renderCell: app => moment(app.createdAt).fromNow(),
    cellClassNames: cx("cell-created"),
  },
]

function ApplicationList() {
  const dispatch = useDispatch()
  const history = useHistory()
  const { pending, applications, error } = useSelector(state => state.applications)

  useEffect(() => {
    dispatch(fetchApplications())
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
        config={tableConfig}
        dataSource={applications}
        onRowClick={app => history.push(`/apps/${app.id}`)}
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
