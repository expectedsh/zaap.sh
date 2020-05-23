import React, { useEffect, useMemo } from "react"
import { toast } from "react-toastify"
import { useHistory } from "react-router"
import { Field, Form } from "react-final-form"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import { createApplication } from "~/store/applications/actions"
import { fetchRunners } from "~/store/runners/actions"
import TextField from "~/oldcomponents/TextField"
import Button from "~/oldcomponents/Button"
import Header from "~/oldcomponents/Header"
import SelectField from "~/oldcomponents/SelectField"
import style from "./ApplicationNew.module.scss"

const cx = classnames.bind(style)

function ApplicationNew() {
  const history = useHistory()
  const dispatch = useDispatch()
  const { pending: runnersPending, runners } = useSelector(state => state.runners)
  const { pending: isLoading } = useSelector(state => state.applications)
  const runnersOptions = useMemo(
    () => runners?.map(runner => ({ label: runner.name, value: runner.id })),
    [runners],
  )

  useEffect(() => {
    dispatch(fetchRunners())
      .catch(() => toast.error("Could not fetch runners."))
  }, [])

  function validate(values) {
    const errors = {}

    if (!values.name) {
      errors.name = "can't be blank"
    } else if (values.name.length < 3 || values.name.length > 50) {
      errors.name = "the length must be between 3 and 50"
    } else if (!values.name.match(/^[a-z]([-a-z0-9]*[a-z0-9])?$/m)) {
      errors.name = "should only contain letters, numbers, and dashes"
    }

    if (!values.image) {
      errors.image = "can't be blank"
    } else if (!values.image.match(/^(?:.+\/)?([^:]+)(?::.+)?$/m)) {
      errors.description = "invalid image"
    }

    if (!values.runnerId) {
      errors.runnerId = "can't be blank"
    }

    return errors
  }

  function onSubmit(values) {
    return dispatch(createApplication({
      ...values,
      runnerId: values.runnerId,
    }))
      .then(() => {
        toast.success("Application created.")
        history.push("/apps")
      })
      .catch(error => {
        if (error.response.status === 422) {
          return error.data
        }
        toast.error(error.response.statusText)
      })
  }

  return (
    <>
      <Header title="Create new application" centered/>
      <div className={cx("root")}>
        <Form
          validate={validate}
          onSubmit={onSubmit}
          render={({ handleSubmit, pristine }) => (
            <form onSubmit={handleSubmit}>
              <Field component={TextField} name="name" label="Name" placeholder="my-app" required/>
              <Field component={TextField} name="image" label="Image" placeholder="nginx:latest"
                     required/>
              <Field component={SelectField} name="runnerId" label="Runner" required
                     isLoading={runnersPending} options={runnersOptions}/>
              <Button loading={isLoading} disabled={pristine} className="btn btn-success"
                      type="submit">
                Create app
              </Button>
            </form>
          )}
        />
      </div>
    </>
  )
}

export default ApplicationNew
