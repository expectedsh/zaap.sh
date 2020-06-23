import React from 'react'
import PropTypes from 'prop-types'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'
import Callout from '~/components/molecules/Callout'

function AppList({ loading, error, applications }) {
  function renderBody() {
    if (loading) {
      return 'Loading...'
    }
    if (error) {
      return (
        <Callout color="danger">
          {error}
        </Callout>
      )
    }
    return applications
  }

  return (
    <>
      <Header preTitle="Overview" title="Applications">
        <Button outline as="link" to="/apps/new" noMargin>
          New application
        </Button>
      </Header>

      {renderBody()}
    </>
  )
}

AppList.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.object,
  applications: PropTypes.array,
}

export default AppList
