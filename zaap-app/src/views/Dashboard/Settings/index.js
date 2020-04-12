import React, { useMemo } from "react"
import classnames from "classnames/bind"
import style from "./Settings.module.scss"
import { useSelector } from "react-redux"
import Alert from "~/components/Alert"
import { Field, Form } from "react-final-form"
import TextField from "~/components/TextField"
import Button from "~/components/Button"

const cx = classnames.bind(style)

function Settings() {
  const user = useSelector(state => state.user.user)

  let loading = false

  function onSubmit() {

  }

  return (
    <div className={cx('root')}>
      <h1 className={cx('title')}>Settings</h1>
      <div className={cx('content')}>
        <Form
          onSubmit={onSubmit}
          initialValues={user}
          render={({ handleSubmit, submitError }) => (
            <form onSubmit={handleSubmit}>
              {submitError && (
                <Alert className="alert alert-error">{submitError}</Alert>
              )}
              <Field component={TextField} name="firstName" label="First name" />
              <Field component={TextField} type="email" name="email" label="Email" />
              <Field component={TextField} name="schedulerToken" label="Scheduler token" disabled />
              <Button className="btn btn-success" type="submit" loading={loading}>
                Continue
              </Button>
            </form>
          )}
        />
      </div>
    </div>
  )
}

export default Settings
