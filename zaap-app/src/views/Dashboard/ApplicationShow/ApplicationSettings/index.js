import React, { useMemo } from "react"
import FormSection from "~/components/FormSection"
import { Field, Form } from "react-final-form"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import { updateApplication } from "~/store/application/actions"

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
    let errors = {}
    if (!values.name) {
      errors.name = "can't be blank"
    }
    if (!values.image) {
      errors.image = "can't be blank"
    }
    const replicas = parseInt(values.replicas, 10)
    if (isNaN(replicas) || replicas < 0) {
      errors.replicas = 'invalid number of replicas'
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
      render={({ handleSubmit, pristine, invalid }) => (
        <form onSubmit={handleSubmit}>
          <Field component={TextField} name="name" label="Application Name"/>
          <Field component={TextField} name="image" label="Image"/>
          <Field component={TextField} type="number" name="replicas" label="Replicas"/>
          {!pristine && (
            <Button className="btn btn-success" type="submit" disabled={invalid}>
              Save
            </Button>
          )}
        </form>
      )}
    />
  )
}

function EnvironmentForm() {
  const application = useSelector(state => state.application.application)
  const initialValues = useMemo(() =>
    application.deployments
      .find(v => v.id === application.currentDeploymentId)
      .environment
  , [])

  function onSubmit() {

  }

  return (
    <Form
      onSubmit={onSubmit}
      initialValues={initialValues}
      render={({ handleSubmit, pristine }) => (
        <form onSubmit={handleSubmit}>
          {/*<Field component={TextField} name="name" label="Application Name"/>*/}
          {/*<Field component={TextField} name="image" label="Image"/>*/}
          {/*<Field component={TextField} type="number" name="replicas" label="Replicas"/>*/}
          {/*{!pristine && (*/}
          {/*  <Button className="btn btn-success" type="submit" disabled={pristine}>*/}
          {/*    Save*/}
          {/*  </Button>*/}
          {/*)}*/}
        </form>
      )}
    />
  )
}

function ApplicationSettings() {
  return (
    <>
      <FormSection name="General">
        <GeneralForm/>
      </FormSection>
      <FormSection name="Environment">
        <EnvironmentForm/>
      </FormSection>
    </>
  )
}

export default ApplicationSettings
