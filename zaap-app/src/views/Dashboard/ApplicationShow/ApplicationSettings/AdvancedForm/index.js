import React, { useEffect, useMemo } from "react"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { Field, Form } from "react-final-form"
import Button from "~/components/Button"
import SelectField from "~/components/SelectField"
import { fetchClusterRoles } from "~/store/runner/actions"
import { updateApplication } from "~/store/application/actions"

function AdvancedForm() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const { clusterRolesPending, clusterRoles } = useSelector(state => state.runner)
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

  useEffect(() => {
    if (application) {
      dispatch(fetchClusterRoles({ id: application.runnerId }))
        .catch(error => {
          toast.error(error.response.statusText)
        })
    }
  }, [application])

  function onSubmit(values) {
    return dispatch(updateApplication({
      id: application.id,
      roles: values.roles,
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
          <Field component={SelectField} name="roles" label="Roles" required isMulti
                 isLoading={clusterRolesPending} options={clusterRoleOptions}/>
          <Button className="btn btn-success" type="submit" disabled={pristine}>
            Update
          </Button>
        </form>
      )}
    />
  )
}

export default AdvancedForm
