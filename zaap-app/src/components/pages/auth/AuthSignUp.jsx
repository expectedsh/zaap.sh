import React from 'react'
import PropTypes from 'prop-types'
import { Field, Form } from 'react-final-form'
import AuthContainer from '~/components/organisms/AuthContainer'
import Input from '~/components/molecules/Input'
import Button from '~/components/atoms/Button'
import Callout from '~/components/molecules/Callout'
import Link from '~/components/atoms/Link'

function AuthSignIn({ onSubmit }) {
  const alternative = <>Already have an account? <Link to="/sign_in">Sign In</Link>.</>

  return (
    <AuthContainer title="Sign Up" alternative={alternative}>
      <Form
        onSubmit={onSubmit}
        render={({ handleSubmit, submitting, submitError }) => (
          <form onSubmit={handleSubmit}>
            {submitError && (
              <Callout block color="danger">{submitError}</Callout>
            )}
            <Field component={Input} large name="firstName" placeholder="First name" />
            <Field component={Input} large type="email" name="email" placeholder="Email" />
            <Field component={Input} large type="password" name="password" placeholder="Password" />
            <Button color="success" size="large" type="submit" loading={submitting} block noMargin>
              Continue
            </Button>
          </form>
        )}
      />
    </AuthContainer>
  )
}

AuthSignIn.propTypes = {
  onSubmit: PropTypes.func.isRequired,
}

export default AuthSignIn
