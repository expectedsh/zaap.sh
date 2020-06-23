import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import { Field, Form } from 'react-final-form'
import Input from '~/components/molecules/Input'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'
import Container from '~/components/atoms/Container'

function AppNew({ isRunnerLoading, runners, onSubmit }) {
  const runnersOptions = useMemo(
    () => runners?.map((runner) => ({ label: runner.name, value: runner.id })),
    [runners],
  )

  function validate(values) {
    const errors = {}

    if (!values.name) {
      errors.name = "can't be blank"
    } else if (values.name.length < 3 || values.name.length > 50) {
      errors.name = 'the length must be between 3 and 50'
    } else if (!values.name.match(/^[a-z]([-a-z0-9]*[a-z0-9])?$/m)) {
      errors.name = 'should only contain letters, numbers, and dashes'
    }

    if (!values.image) {
      errors.image = "can't be blank"
    } else if (!values.image.match(/^(?:.+\/)?([^:]+)(?::.+)?$/m)) {
      errors.description = 'invalid image'
    }

    if (!values.runnerId) {
      errors.runnerId = "can't be blank"
    }

    return errors
  }

  return (
    <>
      <Header title="Create new application" center />
      <Container center>
        <Form
          validate={validate}
          onSubmit={onSubmit}
          render={({ handleSubmit, pristine, submitting }) => (
            <form onSubmit={handleSubmit}>
              <Field
                component={Input}
                name="name"
                label="Name"
                placeholder="my-app"
                required
              />
              <Field
                component={Input}
                name="image"
                label="Image"
                placeholder="nginx:latest"
                required
              />
              <Field
                component={Input}
                name="runnerId"
                label="Runner"
                required
                isLoading={isRunnerLoading}
                options={runnersOptions}
              />
              <Button
                loading={submitting}
                disabled={pristine}
                color="success"
                type="submit"
              >
                Create app
              </Button>
            </form>
          )}
        />
      </Container>
    </>
  )
}

AppNew.propTypes = {
  isRunnerLoading: PropTypes.bool.isRequired,
  runners: PropTypes.arrayOf(PropTypes.object).isRequired,
  onSubmit: PropTypes.func.isRequired,
}

export default AppNew
