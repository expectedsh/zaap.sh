import React, { useMemo, useState } from "react"
import classnames from "classnames/bind"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { updateApplication } from "~/store/application/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import style from "./EnvironmentForm.module.scss"

const cx = classnames.bind(style)

function EnvironmentForm() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const isLoading = useSelector(state => state.application.updatePending)
  const [key, setKey] = useState("")
  const [value, setValue] = useState("")
  const environment = useMemo(
    () => Object.entries(
      application.deployments
        .find(v => v.id === application.currentDeploymentId)
        .environment
    ),
    [application],
  )

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
      ...Object.fromEntries(environment),
      [key]: undefined,
    })
  }

  function onSubmit(event) {
    event.preventDefault()
    updateEnvironment({
      ...Object.fromEntries(environment),
      [key]: value,
    })
  }

  return (
    <>
      {environment.length ? environment.map(([key, value], index) => (
        <div className={cx("env-var")} key={index}>
          <TextField className={cx("env-var-input")} value={key} disabled/>
          <TextField className={cx("env-var-input")} value={value} disabled/>
          <div className={cx("env-var-actions")}>
            <Button className={cx("btn", "material-icons", "delete-button")}
                    disabled={isLoading} onClick={() => deleteVariable(key)}>
              close
            </Button>
          </div>
        </div>
      )) : (
        <div className={cx('empty-state')}>
          You don't have environment variable
        </div>
      )}
      <form className={cx("env-var")} onSubmit={onSubmit}>
        <TextField className={cx("env-var-input")} placeholder="KEY" value={key}
                   onChange={e => setKey(e.target.value)}/>
        <TextField className={cx("env-var-input")} placeholder="VALUE" value={value}
                   onChange={e => setValue(e.target.value)}/>
        <div className={cx("env-var-actions")}>
          <Button className="btn btn-success" type="submit" disabled={!key} loading={isLoading}>
            Add
          </Button>
        </div>
      </form>
    </>
  )
}

export default EnvironmentForm
