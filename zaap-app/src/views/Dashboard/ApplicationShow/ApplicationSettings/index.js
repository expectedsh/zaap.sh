import React, { useMemo, useState } from "react"
import FormSection from "~/components/FormSection"
import classnames from "classnames/bind"
import { Field, Form } from "react-final-form"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { updateApplication } from "~/store/application/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import style from "./ApplicationSettings.module.scss"

const cx = classnames.bind(style)

function GeneralForm() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const initialValues = useMemo(() => {
    const currentDeployment = application.deployments
      .find(v => v.id === application.currentDeploymentId)
    return {
      ...application,
      ...currentDeployment,
    }
  }, [application])

  function validate(values) {
    const errors = {}

    if (!values.name) {
      errors.name = "can't be blank"
    } else if (values.name.length < 3 || values.name.length > 50) {
      errors.name = "the length must be between 3 and 50"
    } else if (!values.name.match(/^[-a-zA-Z0-9]+$/m)) {
      errors.name = "should only contain letters, numbers, and dashes"
    }

    if (!values.image) {
      errors.image = "can't be blank"
    } else if (!values.image.match(/^(?:.+\/)?([^:]+)(?::.+)?$/m)) {
      errors.description = "invalid image"
    }

    const replicas = parseInt(values.replicas, 10)
    if (isNaN(replicas) || replicas < 0) {
      errors.replicas = "invalid number of replicas"
    }

    return errors
  }

  function onSubmit(values) {
    return dispatch(updateApplication({
      id: application.id,
      name: values.name,
      image: values.image,
      replicas: parseInt(values.replicas, 10),
    }))
      .then(() => {
        toast.success("Application updated.")
      })
      .catch(error => {
        if (error.response.status === 422) {
          return error.data
        }
        toast.error(error.response.statusText)
      })
  }

  return (
    <Form
      validate={validate}
      onSubmit={onSubmit}
      initialValues={initialValues}
      render={({ handleSubmit, pristine }) => (
        <form onSubmit={handleSubmit}>
          <Field component={TextField} name="name" label="Application Name"/>
          <Field component={TextField} name="image" label="Image"/>
          <Field component={TextField} type="number" name="replicas" label="Replicas"/>
          <Button className="btn btn-success" type="submit" disabled={pristine}>
            Update
          </Button>
        </form>
      )}
    />
  )
}

function EnvironmentForm() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const isLoading = useSelector(state => state.application.updatePending)
  const [key, setKey] = useState("")
  const [value, setValue] = useState("")
  const environment = useMemo(() =>
      application.deployments
        .find(v => v.id === application.currentDeploymentId)
        .environment
    , [application])

  function updateEnvironment(environment) {
    return dispatch(updateApplication({
      id: application.id,
      environment,
    }))
      .then(() => {
        toast.success("Application updated.")
        setKey("")
        setValue("")
      })
      .catch(error => {
        toast.error(error.response.statusText)
      })
  }

  function deleteVariable(key) {
    updateEnvironment({
      ...environment,
      [key]: undefined,
    })
  }

  function onSubmit(event) {
    event.preventDefault()
    updateEnvironment({
      ...environment,
      [key]: value,
    })
  }

  return (
    <>
      {Object.entries(environment).map(([key, value], index) => (
        <div className={cx("env-var-line")} key={index}>
          <TextField className={cx("env-var-input")} value={key} disabled/>
          <TextField className={cx("env-var-input")} value={value} disabled/>
          <div className={cx("env-var-button")}>
            <Button className={cx("btn", "material-icons", "env-var-delete-button")}
                    disabled={isLoading} onClick={() => deleteVariable(key)}>
              close
            </Button>
          </div>
        </div>
      ))}
      <form className={cx("env-var-line")} onSubmit={onSubmit}>
        <TextField className={cx("env-var-input")} placeholder="KEY" value={key}
                   onChange={e => setKey(e.target.value)}/>
        <TextField className={cx("env-var-input")} placeholder="VALUE" value={value}
                   onChange={e => setValue(e.target.value)}/>
        <div className={cx("env-var-button")}>
          <Button className="btn btn-success" type="submit" disabled={!key} loading={isLoading}>
            Add
          </Button>
        </div>
      </form>
    </>
  )
}

function ApplicationSettings() {
  return (
    <>
      <FormSection name="General">
        <GeneralForm/>
      </FormSection>
      <FormSection name="Environment"
                   description="Environment variables change the way your application behaves.">
        <EnvironmentForm/>
      </FormSection>
      <FormSection className={cx('section-delete')} name="Delete application"
                   description="Deleting your application is irreversible.">
        <Button className="btn btn-outline-danger">
          Delete application
        </Button>
      </FormSection>
    </>
  )
}

export default ApplicationSettings
