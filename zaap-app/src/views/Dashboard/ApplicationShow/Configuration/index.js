import React, { useMemo } from "react"
import FormSection from "~/components/FormSection"
import { Field, Form } from "react-final-form"
import { useSelector } from "react-redux"
import TextField from "~/components/TextField"
import Button from "~/components/Button"

function GeneralForm() {
  const application = useSelector(state => state.application.application)
  const initialValues = useMemo(() => {
    const currentDeployment = application.deployments
      .find(v => v.id === application.currentDeploymentId)
    return {
      ...application,
      ...currentDeployment,
    }
  }, [])

  function onSubmit() {

  }

  return (
    <Form
      onSubmit={onSubmit}
      initialValues={initialValues}
      render={({ handleSubmit, pristine }) => (
        <form onSubmit={handleSubmit}>
          <Field component={TextField} name="name" label="Application Name"/>
          <Field component={TextField} name="image" label="Image"/>
          <Field component={TextField} type="number" name="replicas" label="Replicas"/>
          {!pristine && (
            <Button className="btn btn-success" type="submit" disabled={pristine}>
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

function Configuration() {
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

export default Configuration
