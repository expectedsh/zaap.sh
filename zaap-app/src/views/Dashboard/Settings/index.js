import React from "react"
import classnames from "classnames/bind"
import { Field, Form } from "react-final-form"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { updateUser } from "~/store/user/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import FormSection from "~/components/FormSection"
import style from "./Settings.module.scss"

const cx = classnames.bind(style)

function ProfileForm() {
  const dispatch = useDispatch()
  const user = useSelector(state => state.user.user)

  function onSubmit(values) {
    return dispatch(updateUser(values))
      .then(() => {
        toast.success('Profile updated.')
      })
      .catch(error => {
        if (error.response.status === 422) {
          return error.data
        }
        toast.error(error.response.statusText);
      })
  }

  return (
    <Form
      onSubmit={onSubmit}
      initialValues={user}
      render={({ handleSubmit, pristine }) => (
        <form onSubmit={handleSubmit}>
          <Field component={TextField} name="firstName" label="First name"/>
          <Field component={TextField} type="email" name="email" label="Email"/>
          <Button className="btn btn-success" type="submit" disabled={pristine}>
            Update
          </Button>
        </form>
      )}
    />
  )
}

function SchedulerForm() {
  const user = useSelector(state => state.user.user)

  return (
    <>
      <TextField name="schedulerToken" label="Scheduler token" value={user.schedulerToken} disabled/>
    </>
  )
}

function Settings() {
  return (
    <div className={cx("root")}>
      <h1 className="header-title">Settings</h1>
      <FormSection
        name="Profile"
        description="Your email address is your identity on Zaap and is used to log in."
      >
        <ProfileForm/>
      </FormSection>
      <FormSection
        name="Scheduler"
        description="Informations about your scheduler."
      >
        <SchedulerForm/>
      </FormSection>
    </div>
  )
}

export default Settings
