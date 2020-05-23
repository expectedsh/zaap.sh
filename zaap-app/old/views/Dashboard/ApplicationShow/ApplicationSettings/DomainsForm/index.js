import React from "react"
import classnames from "classnames/bind"
import { toast } from "react-toastify"
import { useDispatch, useSelector } from "react-redux"
import { Field, Form } from "react-final-form"
import { addDomain, deleteDomain } from "~/store/application/actions"
import TextField from "~/oldcomponents/TextField"
import Button from "~/oldcomponents/Button"
import style from "./DomainsForm.module.scss"

const cx = classnames.bind(style)

function DomainsForm() {
  const dispatch = useDispatch()
  const { defaultDomain, domains } = useSelector(state => state.application.application)
  const isLoading = useSelector(state => state.application.updatePending)

  function validate(values) {
    if (values.domain && !values.domain.match(/^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/gi)) {
      return { domain: "must be a valid domain" }
    }
  }

  function handleResponse(res) {
    return res
      .then(() => {
        toast.success("Application updated.")
      })
      .catch(error => {
        toast.error(error.response.statusText)
      })
  }

  function handleDelete(domain) {
    return handleResponse(dispatch(deleteDomain({ domain })))
  }

  function handleCreate(values) {
    return handleResponse(dispatch(addDomain(values)))
  }

  return (
    <>
      <div className={cx("default-domain")}>
        Default domain:{" "}
        <a href={`https://${defaultDomain}`} target="_blank">
          {defaultDomain}
        </a>
      </div>
      {domains?.length ? domains.map((domain, index) => (
        <div key={index} className={cx("domain")}>
          <TextField className={cx("domain-name")} value={domain} disabled/>
          <div className={cx("domain-action")}>
            <Button className={cx("btn", "material-icons", "delete-button")} disabled={isLoading}
                    onClick={() => handleDelete(domain)}>
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
        onSubmit={handleCreate}
        render={({ handleSubmit, pristine, form }) => (
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
