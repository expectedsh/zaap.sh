import React from 'react'
import PropTypes from 'prop-types'
import Header from '~/components/molecules/Header'
import Container from '~/components/atoms/Container'
import SimpleStateHandler from '~/components/molecules/SimpleStateHandler'
import ProfileSection from './ProfileSection'

function Settings({
  loading, error, user, updateProfile,
}) {
  return (
    <>
      <Header preTitle="Account" title="Settings" />
      <Container>
        <SimpleStateHandler
          loading={loading}
          error={error}
          onSuccess={(
            <>
              <ProfileSection user={user} onSubmit={updateProfile} />
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
  updateProfile: PropTypes.func.isRequired,
}

export default Settings
