import React from "react"
import { FORM_ERROR } from "final-form"
import { Field, Form } from "react-final-form"
import { Link, useHistory } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import TextField from "~/components/AuthTextField"
import Button from "~/components/Button"
import Alert from "~/components/Alert"
import { register } from "~/store/authentication/actions"
import style from "~/views/Auth/Auth.module.scss"

const cx = classnames.bind(style)

function SignUp() {
  const dispatch = useDispatch()
  const history = useHistory()
  const loading = useSelector(state => state.authentication.pending)

  function onSubmit(values) {
    return dispatch(register(values))
      .then(() => history.push("/"))
      .catch(error => {
        if (error.response.status === 422) {
          return error.data
        }
        return { [FORM_ERROR]: error.response.statusText }
      })
  }

  return (
    <>
      <div className={cx("container")}>
        <h1 className={cx("title")}>Sign Up</h1>
        <Form
          onSubmit={onSubmit}
          render={({ handleSubmit, submitError }) => (
            <form onSubmit={handleSubmit}>
              {submitError && (
                <Alert className="alert alert-error">{submitError}</Alert>
              )}
              <Field component={TextField} name="firstName" placeholder="First name" />
              <Field component={TextField} type="email" name="email" placeholder="Email" />
              <Field component={TextField} type="password" name="password" placeholder="Password" />
              <Button className="btn btn-success" type="submit" loading={loading}>
                Continue
              </Button>
            </form>
          )}
        />
      </div>
      <div className={cx("alternative")}>
        Already have an account? <Link to="/sign_in">Sign In</Link>.
      </div>
    </>
  )
}

export default SignUp
