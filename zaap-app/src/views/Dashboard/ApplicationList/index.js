import React, { useEffect } from "react"
import { Link, useHistory } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import classnames from "classnames/bind"
import moment from "moment"
import { fetchApplications } from "~/store/applications/actions"
import { deleteApplication } from "~/store/application/actions"
import Alert from "~/components/Alert"
import Header from "~/components/Header"
import Table from "~/components/Table"
import style from "./ApplicationList.module.scss"
import ApplicationStateBadge from "~/components/ApplicationStateBadge"

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

  function remove(id) {
    dispatch(deleteApplication({ id }))
      .then(() => toast.success("Application deleted."))
      .catch(err => toast.error(err.data?.message || err.response.statusText))
  }

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
