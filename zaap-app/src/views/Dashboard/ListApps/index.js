import React, { useEffect } from 'react'
import { Link } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { fetchApplications } from "~/store/applications/actions"
import ApplicationStateBadge from "~/components/ApplicationStateBadge"
import Alert from "~/components/Alert"

function ListApps() {
  const dispatch = useDispatch()
  const { pending, applications, error } = useSelector(state => state.applications)

  useEffect(() => {
    dispatch(fetchApplications())
  }, [])

  function renderBody() {
    if (pending) {
      return <div>Loading...</div>
    }
    if (error) {
      return <Alert className="alert alert-error" error={error} />
    }
    return applications ? (
      <table className="simple-table">
        <thead>
          <tr>
            <th>Name</th>
            <th>Image</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {applications.map(application => (
            <tr key={application.id}>
              <td>{application.name}</td>
              <td>{application.image}</td>
              <td>
                <ApplicationStateBadge state={application.state}/>
              </td>
              <td>
                <Link to={`/apps/${application.id}`}>View</Link>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    ) : null
  }

  return (
    <div className="container">
      <div className="header">
        <h1 className="header-title">
          Applications
        </h1>
        <Link className="btn btn-success" to="/apps/new">
          New application
        </Link>
      </div>
      {renderBody()}
    </div>
  )
}

export default ListApps
