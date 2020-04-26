import React, { useMemo } from "react"
import { toast } from "react-toastify"
import { useDispatch, useSelector } from "react-redux"
import { Field, Form } from "react-final-form"
import classnames from "classnames/bind"
import { addDomain, removeDomain, updateApplication } from "~/store/application/actions"
import TextField from "~/components/TextField"
import Button from "~/components/Button"
import style from "./DomainsForm.module.scss"

const cx = classnames.bind(style)

function DomainsForm() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const isLoading = useSelector(state => state.application.updatePending)
  const domains = useMemo(() => application.domains, [application])

  function handleResponse(res) {
    return res
      .then(() => {
        toast.success("Application updated.")
      })
      .catch(error => {
        toast.error(error.response.statusText)
      })
  }

  function validate(values) {
    if (values.domain && !values.domain.match(/^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/gi)) {
      return { domain: "must be a valid domain" }
    }
  }

  function onSubmit(values) {
    return handleResponse(dispatch(addDomain(values)))
  }

  return (
    <>
      <div className={cx("default-domain")}>
        Default domain:{" "}
        <a href={`https://${application.defaultDomain}`} target="_blank">
          {application.defaultDomain}
        </a>
      </div>
      {domains?.length ? domains.map((domain, index) => (
        <div key={index} className={cx("domain")}>
          <TextField className={cx("domain-name")} value={domain} disabled/>
          <div className={cx("domain-action")}>
            <Button className={cx("btn", "material-icons", "delete-button")} disabled={isLoading}
              onClick={() => handleResponse(dispatch(removeDomain({ domain })))}>
              close
            </Button>
          </div>
        </div>
      )) : (
        <div className={cx("empty-state")}>
          You don't have custom domains
        </div>
      )}
      <Form
        validate={validate}
        onSubmit={onSubmit}
        render={({ handleSubmit, pristine, form  }) => (
          <form onSubmit={e => handleSubmit(e)?.then(form.reset)} className={cx("domain")}>
            <Field component={TextField} className={cx("domain-name")} name="domain"
                   placeholder="Domain name"/>
            <div className={cx("domain-action")}>
              <Button className="btn btn-success" type="submit" disabled={pristine}
                      loading={isLoading}>
                Add
              </Button>
            </div>
          </form>
        )}
      />
    </>
  )
}

export default DomainsForm
