import React from 'react'
import { toast } from "react-toastify"
import { useHistory } from "react-router"
import { Field, Form } from "react-final-form"
import { useDispatch } from "react-redux"
import classnames from "classnames/bind"
import { createApplication } from "~/store/applications/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import style from "./NewApp.module.scss"

const cx = classnames.bind(style)

function ListApps() {
  const history = useHistory()
  const dispatch = useDispatch()

  function onSubmit(values) {
    return dispatch(createApplication(values))
      .then(() => {
        toast.success('Application created.')
        history.push('/apps')
      })
      .catch(error => {
        if (error.response.status === 422) {
          return error.data
        }
        toast.error(error.response.statusText);
      })
  }

  return (
    <div className={cx('root')}>
      <h1 className="header-title">New application</h1>
      <Form
        onSubmit={onSubmit}
        render={({ handleSubmit }) => (
          <form onSubmit={handleSubmit}>
            <Field component={TextField} name="name" label="Name" placeholder="my-app"/>
            <Field component={TextField} name="image" label="Image" placeholder="nginx:latest"/>
            <Button className="btn btn-success" type="submit">
              Create app
            </Button>
          </form>
        )}
      />
    </div>
  )
}

export default ListApps
