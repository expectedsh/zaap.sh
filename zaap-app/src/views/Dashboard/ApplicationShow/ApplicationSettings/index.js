import React from "react"
import FormSection from "~/components/FormSection"
import classnames from "classnames/bind"
import Button from "~/components/Button"
import GeneralForm from "./GeneralForm"
import EnvironmentForm from "./EnvironmentForm"
import DomainsForm from "./DomainsForm"
import style from "./ApplicationSettings.module.scss"

const cx = classnames.bind(style)

function ApplicationSettings() {
  return (
    <>
      <FormSection name="General">
        <GeneralForm/>
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
        <Button className="btn btn-outline-danger">
          Delete application
        </Button>
      </FormSection>
    </>
  )
}

export default ApplicationSettings
