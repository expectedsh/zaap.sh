import { useDispatch, useSelector } from "react-redux"
import React, { useMemo } from "react"
import { toast } from "react-toastify"
import { Field, Form } from "react-final-form"
import { updateApplication } from "~/store/application/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"

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

export default GeneralForm
