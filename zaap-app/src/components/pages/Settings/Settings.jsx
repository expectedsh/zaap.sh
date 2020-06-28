import React from 'react'
import PropTypes from 'prop-types'
import Header from '~/components/molecules/Header'
import Container from '~/components/atoms/Container'
import SimpleStateHandler from '~/components/molecules/SimpleStateHandler'
import FormSection from '~/components/molecules/FormSection'

function Settings({ loading, error, user }) {
  return (
    <>
      <Header preTitle="Account" title="Settings" />
      <Container>
        <SimpleStateHandler
          loading={loading}
          error={error}
          onSuccess={(
            <>
              <FormSection
                name="Profile"
                description="Your email address is your identity on Zaap and is used to log in."
              >
                hello
              </FormSection>
            </>
          )}
        />
      </Container>
    </>
  )
}

Settings.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.object,
  user: PropTypes.object,
}

export default Settings
