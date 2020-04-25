import React, { useMemo } from "react"
import { toast } from "react-toastify"
import { useHistory } from "react-router"
import { Field, Form } from "react-final-form"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import { createRunner } from "~/store/runners/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import Header from "~/components/Header"
import style from "./RunnerNew.module.scss"

const cx = classnames.bind(style)

function RunnerNew() {
  const history = useHistory()
  const dispatch = useDispatch()
  const isLoading = useSelector(state => state.runners.createPending)
  const initialValues = useMemo(() => ({
    token: Math.random().toString(36).slice(2)
      + Math.random().toString(36).slice(2)
      + Math.random().toString(36).slice(2),
  }), [])

  function validate(values) {
    const errors = {}

    if (!values.name) {
      errors.name = "can't be blank"
    } else if (values.name.length < 3 || values.name.length > 50) {
      errors.name = "the length must be between 3 and 50"
    } else if (!values.name.match(/^[-a-zA-Z0-9]+$/m)) {
      errors.name = "should only contain letters, numbers, and dashes"
    }

    if (values.description?.length > 255) {
      errors.description = "the length must be no more than 255"
    }

    return errors
  }

  function onSubmit(values) {
    return dispatch(createRunner(values))
      .then(() => {
        toast.success("Runner registered.")
        history.push("/runners")
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
      <Header title="Register new runner" centered/>
      <div className={cx("root")}>
        <Form
          initialValues={initialValues}
          validate={validate}
          onSubmit={onSubmit}
          render={({ handleSubmit, pristine }) => (
            <form onSubmit={handleSubmit}>
              <Field component={TextField} name="name" label="Name" placeholder="my-runner"
                     required/>
              <Field component={TextField} name="description" label="Description"/>
              <Field component={TextField} name="url" label="URL" placeholder="localhost:8090"
                     required/>
              <Field component={TextField} name="token" label="Token" disabled required/>
              <Button loading={isLoading} disabled={pristine} className="btn btn-success"
                      type="submit">
                Register runner
              </Button>
            </form>
          )}
        />
      </div>
    </>
  )
}

export default RunnerNew
