import React from "react"
import { useDispatch, useSelector } from "react-redux"
import { toast } from "react-toastify"
import { useHistory } from "react-router"
import classnames from "classnames/bind"
import { deleteApplication } from "~/store/application/actions"
import FormSection from "~/components/FormSection"
import Button from "~/components/Button"
import GeneralForm from "./GeneralForm"
import AdvancedForm from "./AdvancedForm"
import EnvironmentForm from "./EnvironmentForm"
import DomainsForm from "./DomainsForm"
import style from "./ApplicationSettings.module.scss"

const cx = classnames.bind(style)

function ApplicationSettings() {
  const history = useHistory()
  const id = useSelector(state => state.application.application.id)
  const dispatch = useDispatch()

  function handleDelete() {
    dispatch(deleteApplication({ id }))
      .then(() => {
        toast.success("Application deleted.")
        history.push('/apps')
      })
      .catch(error => {
        toast.error(error.response.statusText)
      })
  }

  return (
    <>
      <FormSection name="General">
        <GeneralForm/>
      </FormSection>
      <FormSection name="Advanced Options">
        <AdvancedForm/>
      </FormSection>
      <FormSection
        name="Environment"
        description="Environment variables change the way your application behaves."
      >
        <EnvironmentForm/>
      </FormSection>
      <FormSection
        name="Domains"
        description="You can add custom domains to your application."
      >
        <DomainsForm/>
      </FormSection>
      <FormSection
        className={cx("section-delete")}
        name="Delete application"
        description="Deleting your application is irreversible."
      >
        <Button className="btn btn-outline-danger" onClick={handleDelete}>
          Delete application
        </Button>
      </FormSection>
    </>
  )
}

export default ApplicationSettings
