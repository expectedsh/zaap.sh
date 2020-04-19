import React, { useEffect } from 'react'
import { Link } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { fetchApplications } from "~/store/applications/actions"
import ApplicationStateBadge from "~/components/ApplicationStateBadge"
import Alert from "~/components/Alert"
import { deleteApplication } from "~/store/application/actions"
import { toast } from "react-toastify"
import Header from "~/components/Header"

function ApplicationList() {
  const dispatch = useDispatch()
  const { pending, applications, error } = useSelector(state => state.applications)

  useEffect(() => {
    dispatch(fetchApplications())
  }, [])

  function remove(id) {
    dispatch(deleteApplication({ id }))
      .then(() => toast.success('Application deleted.'))
      .catch(err => toast.error(err.data?.message || err.response.statusText))
  }

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
                <div onClick={() => remove(application.id)}>Delete</div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
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
