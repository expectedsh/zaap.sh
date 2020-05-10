import React, { useEffect, useMemo } from "react"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { Field, Form } from "react-final-form"
import Button from "~/components/Button"
import SelectField from "~/components/SelectField"
import { fetchClusterRoles, fetchImagePullSecrets } from "~/store/runner/actions"
import { updateApplication } from "~/store/application/actions"

function AdvancedForm() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const {
    clusterRolesPending, clusterRoles,
    imagePullSecretsPending, imagePullSecrets,
  } = useSelector(state => state.runner)
  const initialValues = useMemo(() => {
    const currentDeployment = application.deployments
      .find(v => v.id === application.currentDeploymentId)
    return {
      ...application,
      ...currentDeployment,
    }
  }, [application])
  const clusterRoleOptions = useMemo(
    () => clusterRoles?.map(role => ({ label: role.name, value: role.name })),
    [clusterRoles]
  )
  const imagePullSecretOptions = useMemo(
    () => imagePullSecrets?.map(role => ({ label: role.name, value: role.name })),
    [imagePullSecrets]
  )

  useEffect(() => {
    if (application) {
      Promise.all([
        dispatch(fetchClusterRoles({ id: application.runnerId })),
        dispatch(fetchImagePullSecrets({ id: application.runnerId }))
      ])
        .catch(error => {
          toast.error(error.response.statusText)
        })
    }
  }, [application])

  function onSubmit(values) {
    return dispatch(updateApplication({
      id: application.id,
      roles: values.roles,
      imagePullSecrets: values.imagePullSecrets,
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
      onSubmit={onSubmit}
      initialValues={initialValues}
      render={({ handleSubmit, pristine }) => (
        <form onSubmit={handleSubmit}>
          <Field component={SelectField} name="roles" label="Roles" isMulti
                 isLoading={clusterRolesPending} options={clusterRoleOptions}/>
          <Field component={SelectField} name="imagePullSecrets" label="Image Pull Secrets" isMulti
                 isLoading={imagePullSecretsPending} options={imagePullSecretOptions}/>
          <Button className="btn btn-success" type="submit" disabled={pristine}>
            Update
          </Button>
        </form>
      )}
    />
  )
}

export default AdvancedForm
