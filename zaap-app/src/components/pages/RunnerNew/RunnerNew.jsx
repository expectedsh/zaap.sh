import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import { Field, Form } from 'react-final-form'
import Input from '~/components/molecules/Input'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'
import Container from '~/components/atoms/Container'

function RunnerNew({ onSubmit }) {
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
      errors.name = 'the length must be between 3 and 50'
    } else if (!values.name.match(/^[-a-zA-Z0-9]+$/m)) {
      errors.name = 'should only contain letters, numbers, and dashes'
    }

    if (values.description?.length > 255) {
      errors.description = 'the length must be no more than 255'
    }

    return errors
  }

  return (
    <>
      <Header title="Register new runner" center />
      <Container center>
        <Form
          initialValues={initialValues}
          validate={validate}
          onSubmit={onSubmit}
          render={({ handleSubmit, pristine, submitting }) => (
            <form onSubmit={handleSubmit}>
              <Field
                component={Input}
                name="name"
                label="Name"
                placeholder="my-runner"
                required
              />
              <Field
                component={Input}
                name="description"
                label="Description"
              />
              <Field
                component={Input}
                name="url"
                label="URL"
                placeholder="localhost:8090"
                required
              />
              <Field
                component={Input}
                name="token"
                label="Token"
                disabled
                required
              />
              <Button
                loading={submitting}
                disabled={pristine}
                color="success"
                type="submit"
              >
                Register runner
              </Button>
            </form>
          )}
        />
      </Container>
    </>
  )
}

RunnerNew.propTypes = {
  onSubmit: PropTypes.func.isRequired,
}

export default RunnerNew
