import PropTypes from 'prop-types'
import React from 'react'
import { Field, Form } from 'react-final-form'
import Input from '~/components/molecules/Input'
import Button from '~/components/atoms/Button'
import FormSection from '~/components/molecules/FormSection'

function ProfileSection({ user, onSubmit }) {
  return (
    <FormSection
      name="Profile"
      description="Your email address is your identity on Zaap and is used to log in."
    >
      <Form
        onSubmit={onSubmit}
        initialValues={user}
        render={({ handleSubmit, pristine }) => (
          <form onSubmit={handleSubmit}>
            <Field component={Input} name="firstName" label="First name" required />
            <Field component={Input} type="email" name="email" label="Email" required />
            <Button color="success" type="submit" disabled={pristine}>
              Update
            </Button>
          </form>
        )}
      />
    </FormSection>
  )
}

ProfileSection.propTypes = {
  user: PropTypes.object.isRequired,
  onSubmit: PropTypes.func.isRequired,
}

export default ProfileSection
